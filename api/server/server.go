package server

import (
	"deck-of-cards/deck"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	Deck   *deck.Deck
}

// Health check handler - Returns 200 with a JSON payload of {"status":true,"message":"UP","data":"UP"}
func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "UP",
		"data":    "UP",
	})
	if err != nil {
		log.Fatal("FAILED TO ENCODE JSON RESULT")
	}
}

// Logger middleware - Applied to all API calls to show the URL,Duration,IP address of a request
func (s *Server) loggerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Take the time before the request
		t1 := time.Now()
		// Perform the request
		h.ServeHTTP(w, r)
		// Take the time after the request
		t2 := time.Now()
		// Print the log to the screen in this format
		/*
			2021/06/04 21:54:30   [GET]        "/health"   72.152µs        - 172.17.0.1:42212
			DATE       TIME     HTTP METHOD     URL        ELAPSED TIME    - REMOTE ADDRESS
		*/
		log.Printf("[%s] %q %v - %s\n", r.Method, r.URL.String(), t2.Sub(t1), r.RemoteAddr)
	})
}

func NewServer() *Server {
	// Create an instance of the Server struct
	server := &Server{
		Router: mux.NewRouter(),
		Deck:   deck.NewDeck(),
	}

	// Middlewares that applies to all requests
	server.Router.Use(server.loggerMiddleware)

	// A list of all end points
	server.Router.HandleFunc("/health", server.healthHandler).Methods("GET")

	// Deck related APIs
	deckApi := server.Router.PathPrefix("/deck").Subrouter()
	deckApi.HandleFunc("/create", server.Deck.CreateDeckHandler).Methods("POST")
	deckApi.HandleFunc("/open/{deckId}", server.Deck.OpenDeckHandler).Methods("POST")
	deckApi.HandleFunc("/draw/{deckId}", server.Deck.DrawCardHandler).Methods("POST")

	return server
}
