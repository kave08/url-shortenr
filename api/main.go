package main

import (
	"log"
	"os"
	"url-shortenr/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolverURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		return
	}

	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))

}
