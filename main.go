package main

import (
	database "ZOCKET/config"
	"ZOCKET/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connent to db
	database.ConnectDB()
	r := fiber.New()

	// Register routes
	routes.RegisterRoutes(r)

	// Start the server
	r.Listen(":8000")
}
