package main

import (
	"server/database"
	"server/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()
	// Define routes
	app.Get("/register", routes.Register)
	
	app.Listen(":5000")
}