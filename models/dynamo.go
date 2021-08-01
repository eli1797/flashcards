package models

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type cardsDynamo interface {
	PutCardInDeck(userId string, card *Card, deckId string) (*Card, error)
	GetAllCardsFromDeck(userId string, deckId string) ([]*Card, error)
	GetDueCardsFromDeck(userId string, deckId string) ([]*Card, error)
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

func (c cardsDynamoDB) PutCardInDeck(userId string, card *Card, deckId string) (*Card, error) {

	pk := userId + "#DECK#" + deckId
	var sk string
	if card.NextDue == "" {
		sk = "CARD#" + strconv.FormatInt(time.Now().Unix(), 10)
	} else {
		sk = "CARD#" + card.NextDue
	}

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

func (c cardsDynamoDB) GetAllCardsFromDeck(userId string, deckId string) ([]*Card, error) {

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

	cards := make([]*Card, len(res.Items))
	for i, item := range res.Items {
		cards[i] = convertDynamoToCard(item)
	}

	return cards, nil
}

func (c cardsDynamoDB) GetDueCardsFromDeck(userId string, deckId string) ([]*Card, error) {

	pk := userId + "#DECK#" + deckId

	// TODO: could add some sort of buffer here (to get cards also due in the next hour or so)
	// or segment the timestamps to days in put
	sk := "CARD#" + strconv.FormatInt(time.Now().Unix(), 10)

	res, err := c.d.Query(&dynamodb.QueryInput{
		ExpressionAttributeNames:   map[string]*string{
			"#pk": aws.String("userId#TYPE"),
			"#sk": aws.String("TYPE#nextDue"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":pk": {
				S: aws.String(pk),
			},
			":sk": {
				S: aws.String(sk),
			},
		},
		KeyConditionExpression:    aws.String("#pk = :pk AND #sk <= :sk"),
		TableName:        &c.tableName,
	})
	if err != nil {
		return nil, err
	}
	if res.LastEvaluatedKey[c.tableName] != nil {
		return nil, errors.New("need to implement paginated query")
	}

	cards := make([]*Card, len(res.Items))
	for i, item := range res.Items {
		cards[i] = convertDynamoToCard(item)
	}

	return cards, nil
}

func convertDynamoToCard(input map[string]*dynamodb.AttributeValue) *Card {
	// TODO, set deck id

	nextDue := strings.Split(*input["TYPE#nextDue"].S, "#")[1]

	out := &Card{
		ID: *input["cardId"].S,
		Front: *input["front"].S,
		Back: *input["rear"].S,
		NextDue: nextDue,
	}

	return out
}
