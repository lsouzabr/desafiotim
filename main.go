package main

import (
    "encoding/json"
    "log"
    "net/http"
	"strings"
)

type Person struct {
    Letters []string
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



func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/person/create", personCreate)

    err := http.ListenAndServe(":8080", mux)
    log.Fatal(err)
}
