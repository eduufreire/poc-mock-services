package dynamo

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type dynamoService struct {
	client *dynamodb.Client
}

var tableName = "MockServices"

func NewDynamoService(cfg aws.Config) *dynamoService {
	return &dynamoService{
		client: dynamodb.NewFromConfig(cfg),
	}
}

func (d *dynamoService) PutItem(ctx context.Context, item interface{}) error {
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		log.Fatal(err)
	}

	_, err = d.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	})

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (d *dynamoService) GetItemByParams(ctx context.Context, params Params) (*ItemDynamo, error) {
	keyCondition := expression.Key("service").Equal(expression.Value(params.Service)).
		And(expression.Key("statusCode").Equal(expression.Value(params.StatusCode)))

	filter := expression.Name("endpoint").Equal(expression.Value(params.Endpoint))

	expr, _ := expression.NewBuilder().
		WithKeyCondition(keyCondition).
		WithFilter(filter).
		Build()

	queryResult, err := d.client.Query(ctx, &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})
	if err != nil {
		log.Fatal(err)
	}

	if len(queryResult.Items) <= 0 {
		log.Fatal("Item not found")
	}

	result := ItemDynamo{}
	err = attributevalue.UnmarshalMap(queryResult.Items[0], &result)
	if err != nil {
		log.Fatal(err)
	}

	return &result, nil
}
