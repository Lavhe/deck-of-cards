package deck_util

import (
	"testing"
)

func TestForCreateDeck(t *testing.T) {
	// Test for Creating a deck
	DeckUtil := DeckUtil{}

	deck, err := DeckUtil.CreateDeck()
	if err != nil {
		t.Errorf("Expected nil error, Got '%s'", err.Error())
	}
	if deck.DeckId == "" {
		t.Errorf("Every deck must have a deck_id")
	}
	if len(deck.Cards) != NUM_OF_CARDS {
		t.Errorf("Expected '%d' cards, Got '%d'", NUM_OF_CARDS, len(deck.Cards))
	}

	for _, card := range deck.Cards {
		count := 0
		for _, inside := range deck.Cards {
			if inside.String() == card.String() {
				count += 1
			}
		}
		if count != 1 {
			t.Errorf("We have a repeating card '%s'", card)
		}
	}
}

func TestForShuffleDeck(t *testing.T) {
	// Test for shuffling deck
	DeckUtil := DeckUtil{}

	deck, err := DeckUtil.CreateDeck()
	if err != nil {
		t.Errorf("Expected nil error, Got '%s'", err.Error())
	}
	deckCopy := make([]Card, len(deck.Cards))
	copy(deckCopy, deck.Cards)

	DeckUtil.ShuffleDeck(deck.Cards)

	shuffled := false
	for index, card := range deck.Cards {
		if card.String() != deckCopy[index].String() {
			shuffled = true
			break
		}
	}
	if !shuffled {
		t.Errorf("The deck is not shuffled")
	}
}

func TestForFilterCards(t *testing.T) {
	// Test for filtering deck
	DeckUtil := DeckUtil{}

	deck, err := DeckUtil.CreateDeck()
	if err != nil {
		t.Errorf("Expected nil error, Got '%s'", err.Error())
	}

	wantedCards := map[string]bool{
		"AS": true,
		"KD": true,
		"2C": true,
	}

	filteredDeck := DeckUtil.FilterCards(deck.Cards, wantedCards)

	if len(filteredDeck) != len(wantedCards) {
		t.Errorf("Expected '%d' cards, Got '%d'", len(wantedCards), len(filteredDeck))
	}

	for _, card := range filteredDeck {
		if !wantedCards[card.String()] {
			t.Errorf("%s must not be in the deck", card.String())
		}
	}
}
