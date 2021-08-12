package models

type Card struct {
	ID      string `json:"id"`
	DeckId  string
	Front   string
	Back    string
	NextDue string `json:"nextDue"`
}
