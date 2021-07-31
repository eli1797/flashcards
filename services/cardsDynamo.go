package services

import (
	"flashcards/models"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type cardsDynamo interface {
	PutCardInDeck(card *models.Card, deckId *string) (*models.Card, error)
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

func (cDB cardsDynamoDB) PutCardInDeck(card *models.Card, deckId *string) (*models.Card, error) {

	// TODO: @Debt hardcarded user id
	userId := "72145bba-63e4-44ce-8cf9-d0ef772cb846"

	pk := userId + "#DECK#" + *deckId
	sk := "CARD#" + strconv.FormatInt(time.Now().Unix(), 10)

	res, err := cDB.d.UpdateItem(&dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":front": {
				S: aws.String(card.Front),
			},
			":back": {
				S: aws.String(card.Back),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"userId#TYPE":  {S: aws.String(pk)},
			"TYPE#nextDue": {S: aws.String(sk)},
		},
		UpdateExpression: aws.String("SET front = :front, rear = :back"),
		TableName:        &cDB.tableName,
	})
	if err != nil {
		return nil, err
	}

	fmt.Println(res)

	return nil, nil
}
