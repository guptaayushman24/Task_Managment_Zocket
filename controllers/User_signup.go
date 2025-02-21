package controllers

import (
	database "ZOCKET/config"
	"ZOCKET/models"
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GenerateJWT function
var jwtSecret = []byte("abcd@1234")

func GenerateJWT(email string) (string, error) {
	// Define expiration time
	expirationTime := time.Now().Add(time.Hour * 24) // Token valid for 24 hours

	// Create a new token
	claims := jwt.MapClaims{
		"email": email,
		"exp":   expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with secret key
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Singin Route
func Signin(c *fiber.Ctx) error {
	// Parse Request Body
	var body models.Signin
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}

	// Get User Collection
	userCollection := database.GetCollection("Zocket")

	// Find User in Database by Email
	var user models.Signin
	err := userCollection.FindOne(context.TODO(), bson.M{"email": body.Email}).Decode(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "User not found"})
	}

	// Check Password (Assuming you store hashed passwords)
	if user.Password != body.Password { // Ideally, use bcrypt.CompareHashAndPassword()
		return c.Status(401).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	// Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	// Sign Token with Secret
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error generating token"})
	}

	// Return Token to Client
	return c.Status(200).JSON(fiber.Map{
		"message": "Login successful",
		"token":   tokenString,
		"user":    user, // Optional: Send user details
	})
}

// Signup Route
func Userprofile(c *fiber.Ctx) error {
	var user models.User

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request body", "error": err.Error()})
	}

	if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Password == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Please enter all the fields"})
	}

	// Getting the collection
	userCollection := database.GetCollection("Zocket")
	// Checking if email already exist
	err1 := userCollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&user)
	if err1 == nil {
		// If  find an existing user, return an error
		return c.Status(400).JSON(fiber.Map{"message": "Email already exists"})
	} else if err1 != mongo.ErrNoDocuments {
		// Handle any other database error
		return c.Status(500).JSON(fiber.Map{"message": "Database error", "error": err.Error()})
	}
	// Inserting the data into mongodb
	result, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Message": "Data is not inserted"})
	}
	fmt.Println("Inserted user ID:", result.InsertedID)
	// Successfully parsed data
	return c.Status(201).JSON(fiber.Map{"message": "User created successfully", "user": user})

}
