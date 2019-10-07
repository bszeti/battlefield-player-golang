
package main

import (
    "encoding/json"
	"fmt"
	"os"
	"strconv"
	// "io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/bszeti/battlefield-player-golang/services"
)


func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"status\": \"OK\"}")
}

//HitByPlayer gets a Battlefield resource by name  
func hitByPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hitByPlayer := vars["player"]

	log.Printf("HitByPlayer: %s",hitByPlayer)

	currentHealth, err :=  services.HitByPlayer(hitByPlayer)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(currentHealth) 
	}
}

//GetCurrentHealth is
func getCurrentHealth(w http.ResponseWriter, r *http.Request) {
	log.Println("GetCurrentHealth")
	currentHealth :=  services.CurrentHealth()
	
	json.NewEncoder(w).Encode(currentHealth) 
}

func main() {
	fmt.Println("Hello World!")

	maxHealth, err := strconv.Atoi( os.Getenv("BATTLEFIELD_MAX_HEALTH") )
	if (err != nil) {
		panic(err)
	}
	services.SetCurrentHealth(int32(maxHealth))

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", healthHandler)
	router.HandleFunc("/api/hit/{player}", hitByPlayer).Methods("GET")
	router.HandleFunc("/api/status/currenthealth", getCurrentHealth).Methods("GET")

		

	err = http.ListenAndServe(":8080", router)
	if err != nil {
        log.Fatal("ListenAndServe Error: ", err)
    }
}
