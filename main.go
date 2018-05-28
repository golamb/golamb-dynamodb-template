package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func toMapDynamodbAttr(image map[string]events.DynamoDBAttributeValue) map[string]*dynamodb.AttributeValue {
	mapData := make(map[string]*dynamodb.AttributeValue)
	for name, value := range image {
		var dynamodbattr dynamodb.AttributeValue
		data, _ := value.MarshalJSON()
		json.Unmarshal(data, &dynamodbattr)
		mapData[name] = &dynamodbattr
	}
	return mapData
}

type dbData struct {
	Message string
	ID      string
}

// Handle is lambda Handle function
func Handle(ctx context.Context, e events.DynamoDBEvent) {
	var messages []dbData
	for _, record := range e.Records {
		mapData := toMapDynamodbAttr(record.Change.NewImage)
		var message dbData
		dynamodbattribute.UnmarshalMap(mapData, &message)
		messages = append(messages, message)
	}
	log.Println(messages)
}

func main() {
	lambda.Start(Handle)
}
