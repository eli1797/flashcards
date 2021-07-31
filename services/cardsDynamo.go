package services

import (
	"errors"
	"flashcards/models"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type cardsDynamo interface {
	PutCardInDeck(userId string, card *models.Card, deckId string) (*models.Card, error)
	GetAllCardsFromDeck(userId string, deckId string) ([]*models.Card, error)
}

type cardsDynamoDB struct {
	d         *dynamodb.DynamoDB
	tableName string
}

func NewCardsDynamoDB(client *dynamodb.DynamoDB, tableName string) cardsDynamoDB {

	return cardsDynamoDB{
		d:         client,
		tableName: tableName,
	}
}

func (c cardsDynamoDB) PutCardInDeck(userId string, card *models.Card, deckId string) (*models.Card, error) {

	pk := userId + "#DECK#" + deckId
	sk := "CARD#" + strconv.FormatInt(time.Now().Unix(), 10)

	_, err := c.d.UpdateItem(&dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":front": {
				S: aws.String(card.Front),
			},
			":back": {
				S: aws.String(card.Back),
			},
			":cardId": {
				S: aws.String(card.ID),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"userId#TYPE":  {S: aws.String(pk)},
			"TYPE#nextDue": {S: aws.String(sk)},
		},
		UpdateExpression: aws.String("SET cardId = :cardId, front = :front, rear = :back"),
		TableName:        &c.tableName,
	})
	if err != nil {
		return nil, err
	}

	return card, nil
}

func (c cardsDynamoDB) GetAllCardsFromDeck(userId string, deckId string) ([]*models.Card, error) {

	pk := userId + "#DECK#" + deckId

	res, err := c.d.Query(&dynamodb.QueryInput{
		ExpressionAttributeNames:   map[string]*string{
			"#pk": aws.String("userId#TYPE"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":pk": {
				S: aws.String(pk),
			},
		},
		KeyConditionExpression:    aws.String("#pk = :pk"),
		TableName:        &c.tableName,
	})
	if err != nil {
		return nil, err
	}
	if res.LastEvaluatedKey[c.tableName] != nil {
		return nil, errors.New("need to implement paginated query")
	}

	cards := make([]*models.Card, len(res.Items))
	for i, item := range res.Items {
		cards[i] = convertDynamoToCard(item)
	}

	return cards, nil
}

func convertDynamoToCard(input map[string]*dynamodb.AttributeValue) *models.Card {
	out := &models.Card{
		ID: *input["cardId"].S,
		Front: *input["front"].S,
		Back: *input["rear"].S,
	}

	return out
}
