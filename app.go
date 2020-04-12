package main

import (
	"go-http-server/route"
	"log"
	"net/http"
	"os"
)

var logger = log.New(os.Stdout, "go-http-server> ", log.Lmicroseconds)
var router = route.New(logger)

func main() {
	logger.Printf("Starting a server")

	http.HandleFunc("/", router.Route)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
