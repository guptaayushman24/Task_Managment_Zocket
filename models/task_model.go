package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task struct represents a task in the database
type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`                            // Auto-generated MongoDB ID
	Title       string             `bson:"title" json:"title"`                                 // Task Title
	Description string             `bson:"description,omitempty" json:"description,omitempty"` // Optional Task Description
	Status      string             `bson:"status" json:"status"`                               // Status: Pending, In-Progress, Completed
	Priority    string             `bson:"priority,omitempty" json:"priority,omitempty"`       // Low, Medium, High
	Assignee    string             `bson:"assignee,omitempty" json:"assignee,omitempty"`       // Assigned User
	Assignedto  string             `bson:"assignedto,omitempty" json:"assignedto,omitempty"`
	DueDate     time.Time          `bson:"due_date,omitempty" json:"due_date,omitempty"` // Optional Due Date
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`                 // Auto-set timestamp
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`                 // Auto-updated timestamp
}

type TaskAssign struct {
	Assignedto string `bson:"assignedto,omitempty" json:"assignedto,omitempty"`
}

type Deletetask struct {
	Assignedto    string `bson:"assignedto,omitempty" json:"assignedto,omitempty"`
	Titletodelete string `bson:"Titletodelete" json:"Titletodelete"`
}
