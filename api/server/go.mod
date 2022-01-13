module server

go 1.16

replace deck-of-cards/deck => ./deck

require (
	github.com/gorilla/mux v1.8.0
    deck-of-cards/deck v0.0.1
)