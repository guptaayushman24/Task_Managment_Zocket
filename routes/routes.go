package routes

import (
	"ZOCKET/controllers"
	"ZOCKET/middleware"

	"github.com/gofiber/fiber/v2"
)

// Register function to set up routes
func RegisterRoutes(app *fiber.App) {
	taskRoutes := app.Group("/")
	taskRoutes.Post("https://random-leda-sudent1-cacf1171.koyeb.app/Signin", controllers.Signin)
	taskRoutes.Post("https://random-leda-sudent1-cacf1171.koyeb.app/Signup", controllers.Userprofile)
	taskRoutes.Post("https://random-leda-sudent1-cacf1171.koyeb.app/createtask", middleware.JWTMiddleware, controllers.CreateTask)
	taskRoutes.Get("https://random-leda-sudent1-cacf1171.koyeb.app/getalltask", middleware.JWTMiddleware, controllers.GetTasks)
	taskRoutes.Post("https://random-leda-sudent1-cacf1171.koyeb.app/assignedtask", middleware.JWTMiddleware, controllers.AssignedTask)
	taskRoutes.Post("https://random-leda-sudent1-cacf1171.koyeb.app/deletetask", middleware.JWTMiddleware, controllers.Deletethetask)
}
