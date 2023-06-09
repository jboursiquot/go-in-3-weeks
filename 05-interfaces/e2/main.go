package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jboursiquot/go-in-3-weeks/05-interfaces/e2/mypackage"
)

func main() {
	sess, err := session.NewSession(&aws.Config{})
	if err != nil {
		log.Fatalf("failed to create AWS session: %v", err)
	}

	client := dynamodb.New(sess)

	ddbSaver := mypackage.DynamoDBSaver{Client: client}
	if err := ddbSaver.Save(nil, &mypackage.Person{Name: "Johnny"}); err != nil {
		log.Fatalf("failed to save: %v", err)
	}
}
