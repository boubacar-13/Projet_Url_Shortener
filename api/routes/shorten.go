package routes

import(
	"time"
	"os"
	"github.com/boubacar-13/Projet_Url_Shortener/datbase"

)

type request struct {
	URL			string		`json:"url"`
	CustomShort	string		`json:"short"`
	Expiry		time.Duration	`json:"expiry"`
}

type response struct {
	URL			string		`json:"url"`
	CustomShort	string		`json:"short"`
	Expiry		time.Duration	`json:"expiry"`
	RateRemaining	int			`json:"rate_limit"`
	RateLimitReset	time.Duration	`json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Cannot parse JSON"})
	}

	//implement rate limiting
	r2 := database.CreateClient(1)
	defer r2.Close()

	val, err := r2.Get(database.Ctx, c.IP()).Result()
	if err == redis.Nil{
		_ = r2.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		val, _ = r2.Get(database.Ctx, c.IP()).Result()
		valInt, _ := strconv.Atoi(val)

		if valInt <= 0{
			limit, _ := r2.TTL(database.Ctx, c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error":"Rate limit exceeded",
				"rate_limit_reset": limit / time.Nanosecond / time.Minute,
			})
		}
	}

	//check if the input is an actual URL
	if !govalidator.IsURL(body.URL){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid URL"})
	}

	//check for domain error
		if !helpers.RemoveDomainError(body.URL){
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error":"You can't access :)"})
	}

	//enforce https, SSL
	body.URL = helpers.EnforceHTTP(body.URL)

	r2.Decr(database.Ctx, c.IP())
}