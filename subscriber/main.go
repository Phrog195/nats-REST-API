package main

import (
	"fmt"
	"log"

	"github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("test-cluster", "User")

	if err != nil {
		log.Fatal(err)
	}

	defer sc.Close()

	sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	}, stan.StartWithLastReceived())

	// Unsubscribe
	//sub.Unsubscribe()
}
