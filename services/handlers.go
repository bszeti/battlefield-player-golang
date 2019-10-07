package services

import (
	// "flag"

	//"encoding/json"
	// "fmt"
	// "os"
	// "path/filepath"
	// "strings"
	// "errors"

	// "k8s.io/client-go/tools/clientcmd"
	//clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	// "io/ioutil"
	"log"
	"sync/atomic"
	// "net/http"

	
	// "k8s.io/client-go/rest"
	
	

	
)

var currentHealth int32
// var lastHitBy string

//HitByPlayer decreses health and returns current value
func HitByPlayer( player string,) (int32, error){
	log.Printf("HitByPlayer is called: %s",player)

	health := atomic.AddInt32(&currentHealth,-1)
	if (health == 0){
		log.Printf("Killed by player: %s",player)
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
	log.Println("Shutting down")
}