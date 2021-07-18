package models

type User struct {
	Decks []*Deck `json:"decks"`
}
