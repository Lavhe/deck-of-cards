package db

import (
	"encoding/json"
	"os"
	"path"

	"io/ioutil"
)

var (
	DATA_DIR = os.Getenv("DATA_DIR")
)

type DB struct {
	dataDir string
}

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

type DBDeck struct {
	DeckId    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
	Cards     []Card `json:"cards"`
}

// Helper function to get the file path to a deck
func (db *DB) getDeckPath(deckId string) string {
	return path.Join(db.dataDir, deckId+".json")
}

// Saves the given deck to a file
func (db *DB) Save(data DBDeck) error {
	file, _ := json.MarshalIndent(data, "", " ")

	return os.WriteFile(db.getDeckPath(data.DeckId), file, 0666)
}

// Retrieves a deck using the given deckId
func (db *DB) GetDeck(deckId string) (*DBDeck, error) {
	data, err := ioutil.ReadFile(db.getDeckPath(deckId))
	if err != nil {
		return nil, err
	}

	var deck DBDeck
	err = json.Unmarshal(data, &deck)
	if err != nil {
		return nil, err
	}

	return &deck, nil
}

// Updates the content of a given deck
func (db *DB) UpdateDeck(dbDeck DBDeck) error {
	err := os.Remove(db.getDeckPath(dbDeck.DeckId))
	if err != nil {
		return err
	}

	return db.Save(dbDeck)
}

// Constructs the db
func NewDB() *DB {
	dataDir := DATA_DIR
	if len(dataDir) < 1 {
		dataDir = "./data"
	}

	os.Mkdir(dataDir, 0777)
	return &DB{
		dataDir: dataDir,
	}
}
