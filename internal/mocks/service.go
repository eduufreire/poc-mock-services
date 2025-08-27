package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/eduufreire/poc-mock-services/internal/aws/dynamo"
	"github.com/eduufreire/poc-mock-services/internal/cache"
)

func NewMockService(dynamoClient dynamo.DynamoService, cache cache.CacheService) *mockService {
	return &mockService{
		client: dynamoClient,
		cache:  cache,
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

	key := buildKey(queryParams)
	hasCache := m.cache.Get(ctx, key)
	if hasCache != nil {
		var payload any
		json.Unmarshal([]byte(*hasCache), &payload)
		return Response{
			StatusCode: queryParams.StatusCode,
			Payload:    payload,
		}
	}

	item, err := m.client.GetItemByParams(ctx, dynamo.Params(queryParams))
	if err != nil {
		log.Fatal(err)
	}

	parsedPayload, err := json.Marshal(item.Payload)
	if err != nil {
		log.Fatal(err)
	}
	m.cache.Set(ctx, key, parsedPayload)

	return Response{
		StatusCode: item.StatusCode,
		Payload:    item.Payload,
	}
}

func buildKey(params Params) string {
	status := strconv.Itoa(params.StatusCode)
	return fmt.Sprintf("%s:%s:%s",
		strings.ToUpper(params.Service),
		strings.ToUpper(params.Endpoint),
		strings.ToUpper(status),
	)
}
