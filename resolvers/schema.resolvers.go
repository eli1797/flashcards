package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"flashcards/generated"
	"flashcards/models"
)

func (r *queryResolver) Deck(ctx context.Context) (*models.Deck, error) {
	return &models.Deck{}, nil
	// panic(fmt.Errorf("not implemented"))
}

// Query returns main.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
