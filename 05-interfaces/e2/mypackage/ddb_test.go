package mypackage_test

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jboursiquot/go-in-3-weeks/05-interfaces/e2/mypackage"
)

// Note how our testClient does not need to depend on the `mypackage.ddbClient` interface here.
type testClient struct {
	output *dynamodb.PutItemOutput
	err    error
}

// We only need our client to satisfy just the bits we need from the DynamoDB client interface implicitly.
func (c *testClient) PutItemWithContext(ctx aws.Context, input *dynamodb.PutItemInput, opts ...request.Option) (*dynamodb.PutItemOutput, error) {
	return c.output, c.err
}

func TestDynamoDBSaver(t *testing.T) {
	tests := map[string]struct {
		person *mypackage.Person
		err    error // We can even mock out errors to test sad paths
	}{
		"happy path": {
			person: &mypackage.Person{Name: "Johnny"},
		},
		"sad path": {
			person: &mypackage.Person{Name: "Johnny"},
			err:    errors.New("failed to save"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client := &testClient{err: tc.err}
			saver := &mypackage.DynamoDBSaver{Client: client}
			ctx := context.Background()
			if err := saver.Save(ctx, tc.person); err != tc.err {
				t.Errorf("expected %v but got %v", tc.err, err)
			}
		})
	}
}
