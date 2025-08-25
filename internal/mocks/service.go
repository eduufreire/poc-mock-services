package mocks

import (
	"context"
	"log"

	"github.com/eduufreire/poc-mock-services/internal/aws/dynamo"
)

func NewMockService(dynamoClient dynamo.DynamoService) *mockService {
	return &mockService{
		client: dynamoClient,
	}
}

func (m *mockService) SaveMock(ctx context.Context, body Request) {
	item := MapToDynamoSchema(body)
	err := m.client.PutItem(ctx, item)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *mockService) GetMockByParams(ctx context.Context, queryParams Params) Response {
	item, err := m.client.GetItemByParams(ctx, dynamo.Params(queryParams))
	if err != nil {
		log.Fatal(err)
	}

	return Response{
		StatusCode: item.StatusCode,
		Payload:    item.Payload,
	}
}
