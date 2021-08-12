package models

import "flashcards/dynamo"

type Deck struct {
	ID       string  `json:"id"`
	Title    *string `json:"title"`
}

func (d *Deck) DueCards() ([]*Card, error) {
	db := NewCardsDynamoDB(dynamo.New(), "tst-cards")

	// TODO: @Debt hardcoded user
	userId := "72145bba-63e4-44ce-8cf9-d0ef772cb846"

	res, err := db.GetDueCardsFromDeck(userId, d.ID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (d *Deck) AllCards() ([]*Card, error) {
	db := NewCardsDynamoDB(dynamo.New(), "tst-cards")

	// TODO: @Debt hardcoded user
	userId := "72145bba-63e4-44ce-8cf9-d0ef772cb846"

	res, err := db.GetAllCardsFromDeck(userId, d.ID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
