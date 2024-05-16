package main

import(
	"fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/boubacar-13/Projet_Url_Shortener/routes"
)

function setUpRoutes(app *fiber.App){
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

function main(){
	err := godotenv.Load()

	if err!=nil{
		fmt.Println(err)
	}

	app := finer.New()

	app.Use(logger.New())

	setUpRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}