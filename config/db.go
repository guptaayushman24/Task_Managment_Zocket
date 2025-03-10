package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Global variable to hold the DB connection
var DB *mongo.Database

func ConnectDB() {
	// MongoDB connection URI
	uri := "mongodb+srv://guptaayushman24:u3xSISNW8S1w4jcW@cluster0.4b48d.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Ping the database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Could not ping MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")

	// Select database
	DB = client.Database("Zocket")

}

// GetCollection returns a reference to a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}

// Unique index to ensure no duplicate email can be added at the time of the Signup
// Ensure unique index on email field
func EnsureUniqueIndex(collection *mongo.Collection) {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},              // Create index on email field
		Options: options.Index().SetUnique(true), // Make it unique
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatal("Could not create index:", err)
	} else {
		log.Println("✅ Unique index on email created successfully!")
	}
}
