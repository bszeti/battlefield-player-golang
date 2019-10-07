package services

import (
	"os"
	"time"
	"io/ioutil"
	"log"
	"sync/atomic"
)

var currentHealth int32
// var lastHitBy string
var termlogPath = getEnv("TERMINATION_LOG_PATH","/dev/termination-log")

//HitByPlayer decreses health and returns current value
func HitByPlayer( player string,) (int32, error){
	log.Printf("HitByPlayer is called: %s",player)

	health := atomic.AddInt32(&currentHealth,-1)
	if (health == 0){
		log.Printf("Killed by player: %s",player)
		defer shutDown(player)
	}
	return health, nil
}

//CurrentHealth returns current health (0+)
func CurrentHealth() (int32){
	health := currentHealth
	if (health > 0) {
		return health
	} else {
		return 0
	}
}

//SetCurrentHealth simply 
func SetCurrentHealth(health int32){
	currentHealth = health
}



func shutDown(killedBy string) {
	data := []byte(killedBy)
    err := ioutil.WriteFile(termlogPath, data, 0666)
    if (err != nil) {
		log.Println("Failed to write termination log.", termlogPath)
	}
	go func() {
		log.Println("Shutting down.")
		time.Sleep(time.Millisecond * 500)
		os.Exit(0)
	} ()
	log.Println("Done.")
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}