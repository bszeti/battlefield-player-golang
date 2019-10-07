package services

import (
	"time"

	"io/ioutil"
	"log"
	"net/http"

	"math/rand"
)

//Players - list of hostnames other players
var Players []string 
//MyName is this player's name
var MyName string

var client = &http.Client{
	Timeout: time.Second * 1,
  }

//Scheduler - hit an other player
func Scheduler(tick *time.Ticker){
	//First time
	hitRandomPlayer()
	//When ticks
	for range tick.C {
		hitRandomPlayer()
	}
}

//Hit another player
func hitRandomPlayer(){
	player := Players[rand.Intn(len(Players))]
	log.Println("Hitting player",player)

	response, err := client.Get("http://"+player+"/api/hit/"+MyName)
    if err != nil {
        log.Println("The HTTP request failed with error", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        log.Println("Response health:",string(data))
	}
}