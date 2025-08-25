package dynamo

import "context"

type DynamoService interface {
	PutItem(ctx context.Context, item any) error
	// GetItems(ctx context.Context) (*any, error)
	GetItemByParams(ctx context.Context, params Params) (*ItemDynamo, error)
	// DeleteItem(ctx context.Context, uuid string) error
}

type ItemDynamo struct {
	Service    string `dynamodbav:"service"`
	StatusCode int    `dynamodbav:"statusCode"`
	Endpoint   string `dynamodbav:"endpoint"`
	Payload    any    `dynamodbav:"payload"`
}

type Params struct {
	Service    string
	Endpoint   string
	StatusCode int
}
