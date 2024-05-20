package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AkhilSharma90/Redis-GO-URL-Shortener/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

// setup two routes, one for shortening the url the other for resolving the url.
//for example if the short is `4fg`, the user must navigate to `localhost:300/04fg` to redirect to  original URL. The domain can be changes in .env file
func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	app := fiber.New()

	//app.Use(csrf.New())
    // Add CORS middleware
    app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:3001", // React app URL
        AllowMethods: "GET,POST,HEAD,PUT,DELETE,OPTIONS",
        AllowHeaders: "Origin,Content-Type,Accept",
    }))


	app.Use(logger.New())

	setupRoutes(app)
	log.Printf("Starting server on port %s...", os.Getenv("APP_PORT"))
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}