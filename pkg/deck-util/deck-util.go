package deck_util

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type DeckUtil struct {
}

type CardValue string

const (
	A     CardValue = "ACE"
	ONE   CardValue = "1"
	TWO   CardValue = "2"
	THREE CardValue = "3"
	FOUR  CardValue = "4"
	FIVE  CardValue = "5"
	SIX   CardValue = "6"
	SEVEN CardValue = "7"
	EIGHT CardValue = "8"
	NINE  CardValue = "9"
	J     CardValue = "JACK"
	Q     CardValue = "QUEEN"
	K     CardValue = "KING"
)

type CardShape string

const (
	Spades   CardShape = "SPADES"
	Diamonds CardShape = "DIAMONDS"
	Clubs    CardShape = "CLUBS"
	Hearts   CardShape = "HEARTS"
)

const (
	NUM_OF_CARDS = 52
)

type Card struct {
	Value CardValue `json:"value"`
	Shape CardShape `json:"shape"`
}

type Deck struct {
	Cards  []Card `json:"cards"`
	DeckId string `json:"deck_id"`
}

func (c Card) String() string {
	value := string(c.Value)[0:1]
	shape := string(c.Shape)[0:1]

	return fmt.Sprintf("%s%s", value, shape)
}

func (d *DeckUtil) CreateDeck() (*Deck, error) {
	shapes := []CardShape{Spades, Diamonds, Clubs, Hearts}
	values := []CardValue{A, ONE, TWO,
		THREE,
		FOUR,
		FIVE,
		SIX,
		SEVEN,
		EIGHT,
		NINE,
		J,
		Q,
		K}

	cards := []Card{}

	for _, shape := range shapes {
		for _, value := range values {
			cards = append(cards, Card{
				Value: value,
				Shape: shape,
			})
		}
	}

	deck := &Deck{
		DeckId: uuid.NewString(),
		Cards:  cards,
	}

	return deck, nil
}

func (d *DeckUtil) ShuffleDeck(deck []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(a int, b int) {
		deck[a], deck[b] = deck[b], deck[a]
	})
}

func (d *DeckUtil) FilterCards(deck []Card, wantedCards map[string]bool) []Card {
	newDeck := []Card{}

	for _, card := range deck {
		if wantedCards[card.String()] {
			newDeck = append(newDeck, card)
		}
	}

	return newDeck
}

func NewDeckUtil() *DeckUtil {
	return &DeckUtil{}
}
