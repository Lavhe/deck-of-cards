package main

import (
	"deck-of-cards/server"
	"log"
	"net/http"
	"os"
)

var (
	PORT = os.Getenv("PORT")
)

func main() {
	// Init the Server
	server := server.NewServer()
	// default the port to 8000
	if len(PORT) == 0 {
		PORT = "8000"
	}
	log.Println("=> App started and listening on port " + PORT)

	// Create a HTTP listener on the specified port
	log.Fatal(http.ListenAndServe(":"+PORT, server.Router))
}
