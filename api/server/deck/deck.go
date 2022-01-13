package deck

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"deck-of-cards/db"
	"deck-of-cards/deck_util"

	"github.com/gorilla/mux"
)

type Deck struct {
	deckUtil *deck_util.DeckUtil
	db       *db.DB
}

func (d *Deck) convertToDBDeck(deck *deck_util.Deck, shuffle bool) db.DBDeck {

	dbDeckCards := []db.Card{}
	for _, card := range deck.Cards {
		dbDeckCards = append(dbDeckCards, db.Card{
			Value: string(card.Value),
			Suit:  string(card.Shape),
			Code:  card.String(),
		})
	}

	dbDeck := db.DBDeck{
		DeckId:    deck.DeckId,
		Shuffled:  shuffle,
		Remaining: len(deck.Cards),
		Cards:     dbDeckCards,
	}

	return dbDeck
}

// Create deck handler - Returns 200 with a JSON payload of {"status":true,"message":"Deck created","data":[]}
func (d *Deck) CreateDeckHandler(w http.ResponseWriter, r *http.Request) {

	deck, err := d.deckUtil.CreateDeck()

	query := r.URL.Query()

	// Filter only wanted cards
	cards := query.Get("cards")
	if len(cards) > 0 {
		wantedCards := strings.Split(cards, ",")
		wantedMap := map[string]bool{}
		for _, card := range wantedCards {
			wantedMap[card] = true
		}
		deck.Cards = d.deckUtil.FilterCards(deck.Cards, wantedMap)
	}

	// Shuffle the cards
	shuffleString := query.Get("shuffle")
	shuffle := false
	if len(shuffleString) > 0 && strings.ToLower(shuffleString) == "true" {
		shuffle = true
		d.deckUtil.ShuffleDeck(deck.Cards)
	}

	dbDeck := d.convertToDBDeck(deck, shuffle)

	err = d.db.Save(dbDeck)
	if err != nil {
		log.Fatal("Unable to save deck", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"deck_id":   dbDeck.DeckId,
		"shuffled":  dbDeck.Shuffled,
		"remaining": dbDeck.Remaining,
	})
	if err != nil {
		log.Fatal("FAILED TO ENCODE JSON RESULT", err)
	}
}

// Open deck handler - Returns 200 with a JSON payload of TODO {"status":true,"message":"","data":[{name:"c","size":4023}]}
func (d *Deck) OpenDeckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	deckId := mux.Vars(r)["deckId"]
	var err error
	if len(deckId) == 0 {
		w.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "A valid deck id is required!",
		})
		return
	}

	deck, err := d.db.GetDeck(deckId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Could not find a deck from the provided Id",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(deck)
	if err != nil {
		log.Fatal("FAILED TO ENCODE JSON RESULT")
	}
}

// Draw cr handler - Returns 200 with a JSON payload of {"status":true,"message":"files","data":[{name:"c","size":4023}]}
func (d *Deck) DrawCardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	deckId := mux.Vars(r)["deckId"]

	query := r.URL.Query()
	count, err := strconv.Atoi(query.Get("count"))
	if err != nil || count < 1 {
		count = 1
	}

	if len(deckId) == 0 {
		w.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "A valid deck id is required!",
		})
		return
	}

	deck, err := d.db.GetDeck(deckId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Could not find a deck from the provided Id",
		})
		return
	}

	if count > len(deck.Cards) {
		w.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"error": fmt.Sprintf("%s %d %s %d", "You can not draw", count, "cards, The deck only has", len(deck.Cards)),
		})
		return
	}
	drawn, remaining := deck.Cards[len(deck.Cards)-count:], deck.Cards[:len(deck.Cards)-count]
	deck.Cards = remaining
	deck.Remaining = len(remaining)

	fmt.Println(deck.Cards)
	err = d.db.UpdateDeck(*deck)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Failed to update the deck" + err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"cards": drawn,
	})
	if err != nil {
		log.Fatal("FAILED TO ENCODE JSON RESULT")
	}
}

func NewDeck() *Deck {
	return &Deck{
		deckUtil: deck_util.NewDeckUtil(),
		db:       db.NewDB(),
	}
}
