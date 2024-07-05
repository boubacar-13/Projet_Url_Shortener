package routes

import (
    "github.com/boubacar-13/Projet_Url_Shortener/database"
    "github.com/gofiber/fiber/v2"
    "github.com/go-redis/redis/v8"
    "log"
)

type StatsResponse struct {
    TotalURLs int64            `json:"total_urls"`
    Clicks    map[string]int64 `json:"clicks"`
}

func GetStats(c *fiber.Ctx) error {
    r := database.CreateClient(1)
    defer r.Close()

    totalURLs, err := r.Get(database.Ctx, "global:total_urls").Int64()
    if err == redis.Nil {
        totalURLs = 0
    } else if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot connect to DB",
        })
    }

    // Get the list of all short URLs
    keys, err := r.Keys(database.Ctx, "clicks:*").Result()
    if err != nil {
        log.Println("Error fetching keys from Redis:", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot connect to DB",
        })
    }

    clicks := make(map[string]int64)
    for _, key := range keys {
        shortURL := key[len("clicks:"):]
        clickCount, err := r.Get(database.Ctx, key).Int64()
        if err == nil {
            clicks[shortURL] = clickCount
        }
    }

    return c.Status(fiber.StatusOK).JSON(StatsResponse{
        TotalURLs: totalURLs,
        Clicks:    clicks,
    })
}
