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

	// Ensuting the unique index on mail

	collection := database.GetCollection("users")
	database.EnsureUniqueIndex(collection)

	// Start the server
	r.Listen(":8000")
}
