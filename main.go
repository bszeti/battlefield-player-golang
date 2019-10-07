package main

import (
    "encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
	"log"
	"net/http"
	"strings"

	"math/rand"

	"github.com/gorilla/mux"
	"github.com/bszeti/battlefield-player-golang/services"
)

//Simply health endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"status\": \"OK\"}")
}

//HitByPlayer takes a hit from another player
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

//GetCurrentHealth return current health
func getCurrentHealth(w http.ResponseWriter, r *http.Request) {
	log.Println("GetCurrentHealth")
	currentHealth :=  services.CurrentHealth()
	
	json.NewEncoder(w).Encode(currentHealth) 
}


var hitPeriod int

func main() {
	fmt.Println("Hello World!")

	//Config - maxHealth
	maxHealth, err := strconv.Atoi( os.Getenv("BATTLEFIELD_MAX_HEALTH") )
	if (err != nil) {
		log.Fatal("Wrong BATTLEFIELD_MAX_HEALTH")
	}
	services.SetCurrentHealth(int32(maxHealth))

	//Config - Hit period
	hitPeriod, err = strconv.Atoi( os.Getenv("BATTLEFIELD_HIT_PERIOD_MS") )
	if (err != nil) {
		log.Fatal("Wrong BATTLEFIELD_HIT_PERIOD_MS")
	}

	//Config - My name
	services.MyName = os.Getenv("BATTLEFIELD_PLAYER_NAME")
	if len(services.MyName)==0 {
		log.Fatal("Wrong BATTLEFIELD_PLAYER_NAME")
	}

	//Config - Other players
	if os.Getenv("BATTLEFIELD_PLAYER_URLS")=="" {
		log.Fatal("Wrong BATTLEFIELD_PLAYER_URLS")
	}
	services.Players = strings.Split(os.Getenv("BATTLEFIELD_PLAYER_URLS"), ",")

	
	rand.Seed(time.Now().Unix()) // init random generator
	//Start scheduled service
	tick := time.NewTicker( time.Duration(hitPeriod) * time.Millisecond )
	go services.Scheduler(tick)	


	//Run http server 
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", healthHandler)
	router.HandleFunc("/api/hit/{player}", hitByPlayer).Methods("GET")
	router.HandleFunc("/api/status/currenthealth", getCurrentHealth).Methods("GET")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
        log.Fatal("ListenAndServe Error: ", err)
	}
	
}
