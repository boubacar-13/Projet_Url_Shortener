package routes

import (
	"os"
	"strconv"
	"time"
	"log"

	"github.com/boubacar-13/Projet_Url_Shortener/database"
	"github.com/boubacar-13/Projet_Url_Shortener/helpers"
	"github.com/asaskevich/govalidator"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {

	// check for the incoming request body
	body := new(request)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	// implement rate limiting: user able to send 10 requests every 30min
	r2 := database.CreateClient(1)
	defer r2.Close()

	log.Println("Getting value from Redis...")

	val, err := r2.Get(database.Ctx, c.IP()).Result()
	
	if err == redis.Nil {
		log.Println("Value not found in Redis. Setting new value...")
		_ = r2.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()

	} else {
		log.Println("Value found in Redis. Checking quota...")
		
		val, _ = r2.Get(database.Ctx, c.IP()).Result()
		valInt, _ := strconv.Atoi(val)
		
		if valInt <= 0 {
			log.Println("Rate limit exceeded. Returning error...")
			limit, _ := r2.TTL(database.Ctx, c.IP()).Result()
			
			log.Println("Rate limit reset:", limit)
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error":            "Rate limit exceeded",
				"rate_limit_reset": limit / time.Nanosecond / time.Minute,
			})
		}
	}

	// check if the input is an actual URL
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid URL",
		})
	}

	// check for the domain error(dont short `localhost:3000` = inifite loop)
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "haha... nice try",
		})
	}

	// enforce https
	body.URL = helpers.EnforceHTTP(body.URL)

	// check if the user has provided any custom short urls
	var id string
	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
	}

	r := database.CreateClient(0)
	defer r.Close()

	val, _ = r.Get(database.Ctx, id).Result()

	// check if the user provided short is already in use
	if val != "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "URL short already in use",
		})
	}
	if body.Expiry == 0 {
		body.Expiry = 24 // default expiry of 24 hours
	}
	err = r.Set(database.Ctx, id, body.URL, body.Expiry*3600*time.Second).Err()
	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to connect to server",
		})
	}
	// respond with the url, short, expiry in hours, calls remaining and time to reset
	resp := response{
		URL:             body.URL,
		CustomShort:     "",
		Expiry:          body.Expiry,
		XRateRemaining:  10,
		XRateLimitReset: 30,
	}

	r2.Decr(database.Ctx, c.IP())
	val, _ = r2.Get(database.Ctx, c.IP()).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)
	ttl, _ := r2.TTL(database.Ctx, c.IP()).Result()
	resp.XRateLimitReset = ttl / time.Nanosecond / time.Minute
	
	resp.CustomShort = os.Getenv("DOMAIN") + "/" + id

	// Add the URL pair to history list
	historyItem := map[string]interface{}{
	"short_url": os.Getenv("DOMAIN") + "/" + id,
	"long_url":  body.URL,
	"created_at": time.Now().Unix(),
	}
	_ = r.HSet(database.Ctx, "url_history:"+id, historyItem).Err()
	_ = r.LPush(database.Ctx, "url_history_list", id).Err()

	// Increment the global URL counter
	rIncr := database.CreateClient(1)
	defer rIncr.Close()
	_ = rIncr.Incr(database.Ctx, "global:total_urls")

	// Initialize the click counter for the new short URL
	_ = rIncr.Set(database.Ctx, "clicks:"+id, 0, 0).Err()

	return c.Status(fiber.StatusOK).JSON(resp)
}