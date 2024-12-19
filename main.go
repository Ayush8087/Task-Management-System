package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB connection
func connectMongo() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://admin:admin%40academichub@34.55.9.182:27017/admin")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Verify connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}

type Task struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description,omitempty" bson:"description"`
	Status      string             `json:"status" bson:"status"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

// Create Task
func createTask(c echo.Context) error {
	client, err := connectMongo()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("taskDB").Collection("tasks")

	var task Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	result, err := collection.InsertOne(context.TODO(), task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	task.ID = result.InsertedID.(primitive.ObjectID)

	return c.JSON(http.StatusCreated, task)
}

// Get All Tasks
func getTasks(c echo.Context) error {
	client, err := connectMongo()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("taskDB").Collection("tasks")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var tasks []Task
	if err := cursor.All(context.TODO(), &tasks); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tasks)
}

// Get Task by ID
func getTaskByID(c echo.Context) error {
	client, err := connectMongo()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer client.Disconnect(context.TODO())

	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	collection := client.Database("taskDB").Collection("tasks")
	var task Task
	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Task not found")
	}

	return c.JSON(http.StatusOK, task)
}

// Update Task
func updateTask(c echo.Context) error {
	client, err := connectMongo()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer client.Disconnect(context.TODO())

	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	var updatedData Task
	if err := c.Bind(&updatedData); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	updatedData.UpdatedAt = time.Now()

	collection := client.Database("taskDB").Collection("tasks")
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, bson.M{"$set": updatedData})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Task updated successfully")
}

// Delete Task
func deleteTask(c echo.Context) error {
	client, err := connectMongo()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer client.Disconnect(context.TODO())

	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	collection := client.Database("taskDB").Collection("tasks")
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	e.POST("/tasks", createTask)
	e.GET("/tasks", getTasks)
	e.GET("/tasks/:id", getTaskByID)
	e.PUT("/tasks/:id", updateTask)
	e.DELETE("/tasks/:id", deleteTask)

	log.Fatal(e.Start(":8080"))
}
