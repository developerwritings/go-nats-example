package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nats-io/nats.go"
)

func main() {

	nc, _ := nats.Connect(nats.DefaultURL)
	nc.Publish("foo", []byte("Hello World"))
	fmt.Println("fdfdfd")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
