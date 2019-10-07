
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

var players []string 
var myName string
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
	myName = os.Getenv("BATTLEFIELD_PLAYER_NAME")
	if len(myName)==0 {
		log.Fatal("Wrong BATTLEFIELD_PLAYER_NAME")
	}

	//Config - Other players
	if os.Getenv("BATTLEFIELD_PLAYER_URLS")=="" {
		log.Fatal("Wrong BATTLEFIELD_PLAYER_URLS")
	}
	players = strings.Split(os.Getenv("BATTLEFIELD_PLAYER_URLS"), ",")

	
	rand.Seed(time.Now().Unix()) // init random generator
	//Start scheduled service
	tick := time.NewTicker( time.Duration(hitPeriod) * time.Millisecond )
	go scheduler(tick)	


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

func scheduler(tick *time.Ticker){
	//First time
	hitRandomPlayer()
	//When ticks
	for range tick.C {
		hitRandomPlayer()
	}
}

//Hit another player
func hitRandomPlayer(){
	player := players[rand.Intn(len(players))]
	log.Println("Hitting player",player)
}