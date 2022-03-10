package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)

	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
