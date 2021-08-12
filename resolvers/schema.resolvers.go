package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"flashcards/dynamo"
	"flashcards/generated"
	"flashcards/models"
	"github.com/google/uuid"
	"strconv"
	"time"
)

// Query returns main.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

func (r *queryResolver) Deck(ctx context.Context, deckId string) (*models.Deck, error) {
	return &models.Deck{
		ID: deckId,
	}, nil
}

func (r *queryResolver) Card(ctx context.Context, deckId string, id string) (*models.Card, error) {
	return &models.Card{
		DeckId: deckId,
	}, nil
}

func (r *queryResolver) User(ctx context.Context) (*models.User, error) {
	// TODO: @HARD hardcoded user name
	return &models.User{
		Id: "72145bba-63e4-44ce-8cf9-d0ef772cb846",
	}, nil
}

// Mutations

// Mutation returns main.Mutation implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) PutCardInDeck(ctx context.Context, deckID string, card *models.CardInput) (*models.Card, error) {

	// TODO: @Debt hardcoded user id, probably need a function that gets the current user from context
	userId := "72145bba-63e4-44ce-8cf9-d0ef772cb846"

	// TODO: @Debt probably want to make a uuid type
	if card.ID == nil {
		newId := uuid.New().String()
		card.ID = &newId
	}

	if card.NextDue == nil {
		*card.NextDue = strconv.FormatInt(time.Now().Unix(), 10)
	}

	result := &models.Card{
		ID:    *card.ID,
		Front: *card.Front,
		Back:  *card.Back,
		NextDue: *card.NextDue,
	}

	db := models.NewCardsDynamoDB(dynamo.New(), "tst-cards")

	res, err := db.PutCardInDeck(userId, result, deckID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *mutationResolver) EditCardNextDue(ctx context.Context, deckID string, cardID string, numMins int) (*string, error) {

	// TODO: @Debt hardcoded user id, probably need a function that gets the current user from context
	userId := "72145bba-63e4-44ce-8cf9-d0ef772cb846"

	db := models.NewCardsDynamoDB(dynamo.New(), "tst-cards")

	/*
	Could change nextDue representation in dynamo to a number. Then could do update adds and not have to call
	However, that's not idempotent.
	 */
	card, err := db.GetCardFromDeckById(userId, deckID, cardID)
	if err != nil {
		return nil, err
	}

	// determine next due time
	curDueTime, err := strconv.Atoi(card.NextDue)
	if err != nil {
		return nil, err
	}

	secondsToAdd := numMins * 60
	newDueTime := curDueTime + secondsToAdd
	newDueTimeStr := strconv.Itoa(newDueTime)

	err = db.MoveCardNextDue(userId, deckID, cardID, newDueTimeStr)
	if err != nil {
		return nil, err
	}

	return &newDueTimeStr, nil
}
