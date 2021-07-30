package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"flashcards/generated"
	"flashcards/models"
)

// Query returns main.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

func (r *queryResolver) Deck(ctx context.Context, id string) (*models.Deck, error) {
	return &models.Deck{}, nil
}

func (r *queryResolver) Card(ctx context.Context, deckId string, id string) (*models.Card, error) {
	return &models.Card{}, nil
}

func (r *queryResolver) User(ctx context.Context) (*models.User, error) {
	// TODO: @HARD hardcoded user name
	return &models.User{
		Id: "72145bba-63e4-44ce-8cf9-d0ef772cb846",
	}, nil
}
