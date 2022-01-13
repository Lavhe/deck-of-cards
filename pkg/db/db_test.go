package db

import (
	"os"
	"testing"

	"github.com/google/uuid"
)

/**
* This tests the following
* - Saves a deck
* - Gets the deck by deckId
* - Draws n cards from the deck
* - Retrieves the deck again
* - Validates if all the conditions are meet
**/
func TestForSave(t *testing.T) {
	db := DB{
		dataDir: os.TempDir(),
	}
	t.Log(db.dataDir)
	deckId := uuid.NewString()
	deck := &DBDeck{
		DeckId:    deckId,
		Shuffled:  false,
		Remaining: 1,
		Cards: []Card{
			{Value: "ACE", Suit: "CLUBS", Code: "AC"},
			{Value: "1", Suit: "CLUBS", Code: "1C"},
			{Value: "3", Suit: "SPADES", Code: "3S"},
			{Value: "7", Suit: "HEARTS", Code: "7H"},
		},
	}

	err := db.Save(*deck)
	if err != nil {
		t.Errorf("Expected nil error, Got '%s'", err.Error())
	}

	newDeck, err := db.GetDeck(deckId)
	if err != nil {
		t.Errorf("Expected nil error, Got '%s'", err.Error())
	}

	if newDeck.DeckId != deck.DeckId {
		t.Errorf("Got the incorrect deck '%s', instead of '%s'", newDeck.DeckId, deck.DeckId)
	}

	if len(newDeck.Cards) != len(deck.Cards) {
		t.Errorf("Got the incorrect number of cards '%d', instead of '%d'", len(newDeck.Cards), len(deck.Cards))
	}

	for index, card := range newDeck.Cards {
		if card.Code != deck.Cards[index].Code {
			t.Errorf("The retrieved deck does not match the saved deck, on '%s' and '%s'", card.Code, deck.Cards[index].Code)
		}
	}

	drawNumber := 3
	deck.Cards = deck.Cards[:len(deck.Cards)-drawNumber]

	err = db.UpdateDeck(*deck)
	if err != nil {
		t.Errorf("Expected nil error, Got '%s'", err.Error())
	}

	updatedDeck, err := db.GetDeck(deckId)

	if updatedDeck.Remaining != len(deck.Cards) {
		t.Errorf("After drawing '%d' cards, We expect the deck to have '%d' cards but instead we got '%d'", drawNumber, len(deck.Cards)-drawNumber, updatedDeck.Remaining)
	}
}
