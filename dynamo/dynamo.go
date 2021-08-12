package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func New() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			MaxRetries: aws.Int(-1),
		},
	}))

	// sess.Handlers.Send.PushFront(func(r *request.Request) {
	// 	// Log every request made and its payload
	// 	logger.Printf("Request: %s/%s, Params: %s",
	// 		r.ClientInfo.ServiceName, r.Operation, r.Params)
	// })

	client := dynamodb.New(sess)

	return client
}
