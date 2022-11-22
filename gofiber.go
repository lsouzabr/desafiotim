// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://docs.gofiber.io
// 📝 Github Repository: https://github.com/gofiber/fiber
package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gofiber/fiber/v2"
)

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

// Database settings (insert your own database name and connection URI)
const dbName = "desafiotim"
const mongoURI = "mongodb://admin:admin@localhost:27017/" + dbName

// Employee struct
type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
}

// Connect configures the MongoDB client and initializes the database connection.
// Source: https://www.mongodb.com/blog/post/quick-start-golang--mongodb--starting-and-setup
func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}

func main() {
	// Connect to the database
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	// Create a Fiber app
	app := fiber.New()

	// Get once employee records from MongoDB
	// Docs: https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-read-documents
	app.Get("/desafiotim/:8752", func(c *fiber.Ctx) error {
		// get id by params
		params := c.Params("id")

		_id, err := primitive.ObjectIDFromHex(params)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		filter := bson.D{{"_id", _id}}

		var result Employee

		if err := mg.Db.Collection("desafiotim").FindOne(c.Context(), filter).Decode(&result); err != nil {
			return c.Status(500).SendString("Something went wrong.")
		}

		return c.Status(fiber.StatusOK).JSON(result)
	})






	log.Fatal(app.Listen(":3000"))
}