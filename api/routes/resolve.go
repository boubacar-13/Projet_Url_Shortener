package routes

import (
	"github.com/AkhilSharma90/Redis-GO-URL-Shortener/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"log"
)

// ResolveURL ...
func ResolveURL(c *fiber.Ctx) error {
	
	// get the short from the url
	url := c.Params("url")
	log.Println("Received request to resolve short URL: %s\n", url)


	// query db to find the original URL, if a match is found
	// increment the counter and redirect to the original URL
	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
				log.Println("Short URL not found in Redis: %s\n", url)

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short not found on database",
		})
	} else if err != nil {
		log.Println("Error fetching short URL from Redis: %s, error: %v\n", url, err)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to DB",
		})
	}

	log.Println("Short URL resolved to: %s\n", value)

	// increment the counter
	rInr := database.CreateClient(1)
	defer rInr.Close()
	_ = rInr.Incr(database.Ctx, "counter")

	// redirect to original URL
	return c.Redirect(value, 301)
}