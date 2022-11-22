package main

import (
    "context" // manage multiple requests
    "fmt"     // Println() function
    "os"
    "time"
	"encoding/json"
    // import 'mongo-driver' package libraries
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Fields struct {
    Name  string
    Email string
    Dept  int
}

func main() {
    // Declare host and port options to pass to the Connect() method
    clientOptions := options.Client().ApplyURI("mongodb://admin:admin@localhost:27017")
    //fmt.Println("clientOptions type:", reflect.TypeOf(clientOptions), "\n")

    // Connect to the MongoDB and return Client instance
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        fmt.Println("mongo.Connect() ERROR:", err)
        os.Exit(1)
    }

    // Declare Context type object for managing multiple API requests
    ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

    // Access a MongoDB collection through a database
    col := client.Database("desafiotim").Collection("desafiotim")
    //fmt.Println("Collection type:", reflect.TypeOf(col), "\n")

    // Declare an empty array to store documents returned
    var result Fields

    // Get a MongoDB document using the FindOne() method
    err = col.FindOne(context.TODO(), bson.D{}).Decode(&result)
    if err != nil {
        fmt.Println("FindOne() ERROR:", err)
        os.Exit(1)
    } else {
//        fmt.Println("FindOne() result:", result)
  //      fmt.Println("FindOne() Name:", result.Name)
    //    fmt.Println("FindOne() Dept:", result.Dept)
    }

    // Call the collection's Find() method to return Cursor obj
    // with all of the col's documents
    cursor, err := col.Find(context.TODO(), bson.D{})

    // Find() method raised an error
    if err != nil {
        fmt.Println("Finding all documents ERROR:", err)
        defer cursor.Close(ctx)

    // If the API call was a success
    } else {
        // iterate over docs using Next()
        for cursor.Next(ctx) {

            // declare a result BSON object
            var result bson.M
            err := cursor.Decode(&result)

            // If there is a cursor.Decode error
            if err != nil {
                fmt.Println("cursor.Next() error:", err)
                os.Exit(1)
               
            // If there are no cursor.Decode errors
            } else {
                //fmt.Println("\nresult type:", reflect.TypeOf(result))
                //fmt.Println("result:", result)

				jsonStr, err := json.Marshal(result)
				if err != nil {
					fmt.Printf("Error: %s", err.Error())
				} else {
					fmt.Println(string(jsonStr))
				}				


            }
        }
    }
}