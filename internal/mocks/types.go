package mocks

import (
	"github.com/eduufreire/poc-mock-services/internal/aws/dynamo"
)

type mockService struct {
	client dynamo.DynamoService
}

type mockHandler struct {
	service *mockService
}

type Request struct {
	Endpoint   string
	Service    string
	StatusCode int
	Payload    map[string]any
}

type Response struct {
	StatusCode int
	Payload    any
}

type Params struct {
	Service    string
	Endpoint   string
	StatusCode int
}
