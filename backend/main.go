package main

import (
	"backend/database"
	"log"

	"backend/users"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	database.SetupDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	users.SetupRoutes(app)

	err := app.Listen(":3333")

	if err != nil {
		log.Fatal("Failed to start the server: ", err)
	}
}
