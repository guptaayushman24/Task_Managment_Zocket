package controllers

import (
	"context"
	"time"

	database "ZOCKET/config"
	"ZOCKET/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateTask - Inserts a new task into MongoDB
func CreateTask(c *fiber.Ctx) error {
	var task models.Task

	// Parse request body into task struct
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request body", "error": err.Error()})
	}

	// Set auto-generated fields
	task.ID = primitive.NewObjectID()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	// Get collection
	taskCollection := database.GetCollection("tasks")

	// Insert into MongoDB
	_, err := taskCollection.InsertOne(context.TODO(), task)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to create task", "error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Task created successfully", "task": task})
}

// Retriving all the task from the Mongodb
func GetTasks(c *fiber.Ctx) error {
	taskCollection := database.GetCollection("tasks")
	// Define projection to include only required fields
	projection := bson.M{
		"title":       1,
		"description": 1,
		"status":      1,
		"priority":    1,
		"_id":         0, // Exclude MongoDB _id field
	}

	// Find all tasks with the projection
	cursor, err := taskCollection.Find(context.TODO(), bson.M{}, &options.FindOptions{
		Projection: projection,
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to retrieve tasks", "error": err.Error()})
	}

	var tasks []models.Task
	if err = cursor.All(context.TODO(), &tasks); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to parse tasks", "error": err.Error()})
	}

	return c.JSON(tasks)
}

// Task assigned to the user
func AssignedTask(c *fiber.Ctx) error {
	var taskassignedto models.TaskAssign

	// Parse request body into task struct
	if err := c.BodyParser(&taskassignedto); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request body", "error": err.Error()})
	}
	taskCollection := database.GetCollection("tasks")
	// Define projection to include only required fields
	projection := bson.M{
		"title":       1,
		"description": 1,
		"status":      1,
		"priority":    1,
		"_id":         0, // Exclude MongoDB _id field
	}

	// Find all tasks with the projection
	cursor, err := taskCollection.Find(context.TODO(), bson.M{"assignedto": taskassignedto.Assignedto}, &options.FindOptions{
		Projection: projection,
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to retrieve tasks", "error": err.Error()})

	}

	var tasks []models.Task
	if err = cursor.All(context.TODO(), &tasks); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to parse tasks", "error": err.Error()})
	}

	return c.JSON(tasks)

}

// Deleting the task which are done
func Deletethetask(c *fiber.Ctx) error {
	taskCollection := database.GetCollection("tasks")
	var tasktodelete models.Deletetask
	// Parse request body into task struct
	if err := c.BodyParser(&tasktodelete); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request body", "error": err.Error()})
	}

	// Delete one document matching both email and title
	filter := bson.M{"assignedto": tasktodelete.Assignedto, "title": tasktodelete.Titletodelete}
	result, err := taskCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to delete task", "error": err.Error()})
	}

	// Check if a document was actually deleted
	if result.DeletedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "No matching task found"})
	}

	return c.JSON(fiber.Map{"message": "Task deleted successfully"})

}
