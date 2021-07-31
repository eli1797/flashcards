package models

type Deck struct {
	ID       string  `json:"id"`
	Title    *string `json:"title"`
	dueCards []*Card
}

func (d *Deck) DueCards() ([]*Card, error) {
	return nil, nil
}

func (d *Deck) AllCards() ([]*Card, error) {
	return nil, nil
}
