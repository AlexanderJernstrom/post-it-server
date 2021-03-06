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
	app.Post("/register", routes.Register)
	app.Post("/login", routes.Login)
	
	app.Listen(":5000")
}