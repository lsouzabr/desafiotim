package main

import (
    "encoding/json"
    "log"
    "net/http"
	"strings"
    "context" 
    "fmt"     
    "os"
    "time"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"	
	"go.mongodb.org/mongo-driver/bson/primitive"

)

type Person struct {
    Letters []string
}

type Fields struct {
    Name  string
    Email string
    Dept  int
}

func valida(w http.ResponseWriter, r *http.Request) {
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
	
		}
		if(len(acumula)>0){
			w.Header().Set("Content-Type", "application/json")

			jsonStr := `[{"is_valid":true}]`
			w.Write([]byte(jsonStr))
			updateStatus(jsonStr)
		} else{
			w.Header().Set("Content-Type", "application/json")

			jsonStr := `[{"is_valid":false}]`
			updateStatus(jsonStr)
			w.Write([]byte(jsonStr))
		}
}

func updateStatus(dados string){


    clientOptions := options.Client().ApplyURI("mongodb://admin:admin@localhost:27017")

    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        fmt.Println("mongo.Connect() ERROR:", err)
        os.Exit(1)
    }

    ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

    col := client.Database("desafiotim").Collection("desafiotim")

    var result Fields

    err = col.FindOne(context.TODO(), bson.D{}).Decode(&result)
    if err != nil {
        fmt.Println("FindOne() ERROR:", err)
        os.Exit(1)
    } 

    cursor, err := col.Find(context.TODO(), bson.D{})


    if err != nil {
        fmt.Println("Finding all documents ERROR:", err)
        defer cursor.Close(ctx)

    } else {
        for cursor.Next(ctx) {

            var result bson.M
            err := cursor.Decode(&result)

            if err != nil {
                fmt.Println("cursor.Next() error:", err)
                os.Exit(1)
               
            } else {
				jsonStr, err := json.Marshal(result)


			type MyJsonName struct {
				ID struct {
					id string `json:"$_id"`
				} `json:"_id"`
				CountInvalid int64   `json:"count_invalid"`
				CountValid   float64 `json:"count_valid"`
				Ratio        float64 `json:"ratio"`
			}

			// web server

			type Foo struct {
				Number int    `json:"number"`
				Title  string `json:"title"`
			}


			var mJsonName MyJsonName
			json.Unmarshal(jsonStr, &mJsonName)
			fmt.Printf("Error: %s", jsonStr)



				coll := client.Database("desafiotim").Collection("desafiotim")
				id, _ := primitive.ObjectIDFromHex("637d01b595bc98257139a3f8")
				filter := bson.D{{"_id", id}}
				update := bson.D{{"$set", bson.D{{"count_valid", 5.4}}}}
				result, err := coll.UpdateOne(context.TODO(), filter, update)
				fmt.Println("", result)
				




				if err != nil {
					panic(err)
				}


				if err != nil {
					fmt.Printf("Error: %s", err.Error())
				} else {
					fmt.Println(string(jsonStr))
				}				
            }

        }
    }
}

func carrega(w http.ResponseWriter, r *http.Request) {
    clientOptions := options.Client().ApplyURI("mongodb://admin:admin@my-mongodb:27017")
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        fmt.Println("mongo.Connect() ERROR:", err)
        os.Exit(1)
    }

    ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

    col := client.Database("desafiotim").Collection("desafiotim")

    var result Fields

    err = col.FindOne(context.TODO(), bson.D{}).Decode(&result)
    if err != nil {
        fmt.Println("FindOne() ERROR:", err)
        os.Exit(1)
    } 

    cursor, err := col.Find(context.TODO(), bson.D{})


    if err != nil {
        fmt.Println("Finding all documents ERROR:", err)
        defer cursor.Close(ctx)

    } else {
        for cursor.Next(ctx) {

            var result bson.M
            err := cursor.Decode(&result)

            if err != nil {
                fmt.Println("cursor.Next() error:", err)
                os.Exit(1)
               
            } else {

				w.Header().Set("Content-Type", "application/json")
				jsonStr, err := json.Marshal(result)
				w.Write([]byte(jsonStr))

				//return jsonStr


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
    mux.HandleFunc("/valida", valida)
	mux.HandleFunc("/status", carrega)
    err := http.ListenAndServe(":8081", mux)
    log.Fatal(err)
}
