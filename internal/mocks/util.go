package mocks

import "github.com/eduufreire/poc-mock-services/internal/aws/dynamo"

func MapToDynamoSchema(data Request) dynamo.ItemDynamo {
	return dynamo.ItemDynamo{
		Service:    data.Service,
		StatusCode: data.StatusCode,
		Endpoint:   data.Endpoint,
		Payload:    data.Payload,
	}
}
