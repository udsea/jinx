package main

import (
	"github.com/udsea/jinx-backend/internal/api/handlers"
	"github.com/udsea/jinx-backend/internal/db"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	app := fiber.New()
	db.InitDB()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Jinx Backend 🎧")
	})

	app.Get("/auth/login", handlers.SpotifyLogin)
	app.Get("/auth/callback", handlers.SpotifyCallback)


	log.Fatal(app.Listen(":8080"))
}

