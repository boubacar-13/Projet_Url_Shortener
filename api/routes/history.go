package routes

import (
    "github.com/AkhilSharma90/Redis-GO-URL-Shortener/database"
    "github.com/gofiber/fiber/v2"
    "github.com/go-redis/redis/v8"
    "log"
    "time"
    "strconv"
)

type HistoryResponse struct {
    ShortURL  string `json:"short_url"`
    LongURL   string `json:"long_url"`
    CreatedAt string `json:"created_at"`
}

func GetHistory(c *fiber.Ctx) error {
    r := database.CreateClient(0)
    defer r.Close()

    ids, err := r.LRange(database.Ctx, "url_history_list", 0, -1).Result()
    if err != nil {
        log.Println("Error fetching history list from Redis:", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot connect to DB",
        })
    }

    var history []HistoryResponse
    for _, id := range ids {
        item, err := r.HGetAll(database.Ctx, "url_history:"+id).Result()
        if err == redis.Nil {
            continue
        } else if err != nil {
            log.Println("Error fetching history item from Redis:", err)
            continue
        }

        timestamp, _ := strconv.ParseInt(item["created_at"], 10, 64)
        history = append(history, HistoryResponse{
            ShortURL:  item["short_url"],
            LongURL:   item["long_url"],
            CreatedAt: time.Unix(timestamp, 0).Format(time.RFC3339),
        })
    }

    return c.Status(fiber.StatusOK).JSON(history)
}
