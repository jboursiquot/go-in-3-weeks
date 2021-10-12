package mypackage

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Note the key takeaway here:
// We don't need to mock out the entire DynamoDB client interface, just the methods we need.
// We don't even need to use the interface AWS provides for DynamoDB SDK, again, just the behavior we need.
type ddbClient interface {
	PutItemWithContext(ctx aws.Context, input *dynamodb.PutItemInput, opts ...request.Option) (*dynamodb.PutItemOutput, error)
}

// DynamoDBSaver interacts with DynamoDB.
type DynamoDBSaver struct {
	Client ddbClient
}

// Person captures demographics.
type Person struct {
	Name string
}

// Save saves.
func (s *DynamoDBSaver) Save(ctx context.Context, p *Person) error {
	item, err := dynamodbattribute.MarshalMap(p)
	if err != nil {
		return fmt.Errorf("failed to marshal shoutout for storage: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	}

	_, err = s.Client.PutItemWithContext(ctx, input)

	return err
}
