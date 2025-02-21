package main

import (
	database "ZOCKET/config"
	"ZOCKET/routes"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var clients map[string]*websocket.Conn

// WebSocket endpoint to register clients
func wsHandler(c *websocket.Conn) {
	// For this example, assume the email is sent as the first message.
	_, msg, err := c.ReadMessage()
	if err != nil {
		c.Close()
		return
	}
	userEmail := string(msg)
	clients[userEmail] = c
	fmt.Printf("User %s connected\n", userEmail)

	// Keep connection open
	for {
		if _, _, err := c.ReadMessage(); err != nil {
			delete(clients, userEmail)
			c.Close()
			break
		}
	}
}
func main() {
	// Connent to db
	database.ConnectDB()
	r := fiber.New()

	// Register routes
	routes.RegisterRoutes(r)

	// Ensuting the unique index on mail

	collection := database.GetCollection("users")
	database.EnsureUniqueIndex(collection)
	// WebSocket route
	r.Get("/ws", websocket.New(wsHandler))

	// Start the server
	r.Listen(":8000")
}
