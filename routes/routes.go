package routes

import (
	"ZOCKET/controllers"
	"ZOCKET/middleware"

	"github.com/gofiber/fiber/v2"
)

// Register function to set up routes
func RegisterRoutes(app *fiber.App) {
	taskRoutes := app.Group("/")
	taskRoutes.Post("/Signin", controllers.Signin)
	taskRoutes.Post("/Signup", controllers.Userprofile)
	taskRoutes.Post("/createtask", middleware.JWTMiddleware, controllers.CreateTask)
	taskRoutes.Get("/getalltask", middleware.JWTMiddleware, controllers.GetTasks)
	taskRoutes.Post("/assignedtask", middleware.JWTMiddleware, controllers.AssignedTask)
	taskRoutes.Post("/deletetask", middleware.JWTMiddleware, controllers.Deletethetask)
}
