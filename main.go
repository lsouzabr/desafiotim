package main

import (
    "encoding/json"
    "log"
    "net/http"
	"strings"
    "context" // manage multiple requests
    "fmt"     // Println() function
    "os"
    "time"
	//"reflect"
    // import 'mongo-driver' package libraries
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"	
)

type Person struct {
    Letters []string
}

type Fields struct {
    Name  string
    Email string
    Dept  int
}

func personCreate(w http.ResponseWriter, r *http.Request) {
    // Declare a new Person struct.
    var p Person

	var theArray [6]string
	theArray[0] = "DUHBHB"  
	theArray[1] = "DUBUHD" 
	theArray[2] = "UBUUHU"  
	theArray[3] = "BHBDHH"  
	theArray[4] = "DDDDUB" 
	theArray[5] = "UDBDUH"  

    err := json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	//fmt.Fprintf(w, "\nPalavras enviadas em: %+v", p.Letters)

		acumula := ""

		for i:=0; i<len(p.Letters); i++{
			

	
			
	
			if(strings.Contains(p.Letters[i], "B")){
				acumula += "B"
			}
			if(strings.Contains(p.Letters[i], "U")){
				acumula += "U"
			}		
			if(strings.Contains(p.Letters[i], "D")){
				acumula += "D"
			}		
			if(strings.Contains(p.Letters[i], "H")){
				acumula += "H"
			}		
	
	
			//fmt.Fprintf(w, "\nLetras encontradas: %+v", acumula)
			//fmt.Fprintf(w, "\nPalavras encontradas: %+v", itemExists(theArray,"UDBDUH ==>"))	
	
		}
		if(len(acumula)>0){
			w.Header().Set("Content-Type", "application/json")

			jsonStr := `[{"is_valid":true}]`
		
			w.Write([]byte(jsonStr))
		} else{
			w.Header().Set("Content-Type", "application/json")

			jsonStr := `[{"is_valid":false}]`
		
			w.Write([]byte(jsonStr))
		}

}

func carrega(w http.ResponseWriter, r *http.Request) {
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

				//jsonStr, err := json.Marshal(result)

				w.Header().Set("Content-Type", "application/json")
				jsonStr, err := json.Marshal(result)
				w.Write([]byte(jsonStr))


				if err != nil {
					fmt.Printf("Error: %s", err.Error())
				} else {
					fmt.Println(string(jsonStr))
				}				


            }
        }
    }
}



func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/valida", personCreate)
	mux.HandleFunc("/status", carrega)

    err := http.ListenAndServe(":8080", mux)
    log.Fatal(err)
}
