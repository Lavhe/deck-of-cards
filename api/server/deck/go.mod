module deck

go 1.16

replace deck-of-cards/deck_util => ../../../pkg/deck-util
replace deck-of-cards/db => ../../../pkg/db

require (
    deck-of-cards/deck_util v0.0.1
    deck-of-cards/db v0.0.1
)