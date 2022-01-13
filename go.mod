module deck-of-cards

go 1.16

replace deck-of-cards/server => ./api/server

replace deck-of-cards/deck => ./api/server/deck

replace deck-of-cards/deck_util => ./pkg/deck-util

replace deck-of-cards/db => ./pkg/db

require (
	deck-of-cards/db v0.0.1
	deck-of-cards/deck v0.0.1
	deck-of-cards/deck_util v0.0.1
	deck-of-cards/server v0.0.1
	github.com/google/uuid v1.3.0 // indirect
)
