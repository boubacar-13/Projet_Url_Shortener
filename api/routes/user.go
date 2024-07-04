package routes

import (
    "encoding/json"
    "net/http"
    "github.com/AkhilSharma90/Redis-GO-URL-Shortener/database"
    "github.com/AkhilSharma90/Redis-GO-URL-Shortener/models"
    "github.com/AkhilSharma90/Redis-GO-URL-Shortener/helpers"
    "github.com/AkhilSharma90/Redis-GO-URL-Shortener/utils"
    "github.com/go-redis/redis/v8"
    "github.com/gofiber/fiber/v2"
    "golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {
    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    user.Password = string(hashedPassword)
    user.ID = helpers.GenerateID()

    userJSON, _ := json.Marshal(user)
    database.CreateClient(0).Set(database.Ctx, user.Email, userJSON, 0)

    return c.Status(fiber.StatusCreated).JSON(user)
}

func Login(c *fiber.Ctx) error {
    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
    }

    userJSON, err := database.CreateClient(0).Get(database.Ctx, user.Email).Result()
    if err == redis.Nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
    }

    var storedUser models.User
    json.Unmarshal([]byte(userJSON), &storedUser)

    if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
    }

    token, err := utils.GenerateJWT(storedUser.ID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}
