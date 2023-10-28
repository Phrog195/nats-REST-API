package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/nats-io/stan.go"
)

func main() {
	var order Order
	//Reading JSON from file
	content, err := os.ReadFile("D:/Dev/projects/internship/model.json")
	if err != nil {
		log.Fatal(err)
	}
	//Unmarshalling JSON to GO
	err = json.Unmarshal(content, &order)
	if err != nil {
		log.Fatal("Broken JSON")
	}
	//Creating Channel
	sc, err := stan.Connect("test-cluster", "User")
	if err != nil {
		log.Fatal(err)
	}

	defer sc.Close()

	ackHandler := func(ackedNuid string, err error) {
		if err != nil {
			log.Printf("Warning: error publishing msg id %s: %v\n", ackedNuid, err.Error())
		} else {
			log.Printf("Received ack for msg id %s\n", ackedNuid)
		}
	}

	nuid, err := sc.PublishAsync("foo", content, ackHandler) // returns immediately
	if err != nil {
		log.Printf("Error publishing msg %s: %v\n", nuid, err.Error())
	}
}
