package main

import (
	"net/http"

	"github.com/eduufreire/poc-mock-services/internal/aws"
	"github.com/eduufreire/poc-mock-services/internal/aws/dynamo"
	"github.com/eduufreire/poc-mock-services/internal/mocks"
)

func main() {

	clientDynamo := dynamo.NewDynamoService(aws.Session())

	mockService := mocks.NewMockService(clientDynamo)
	mockHandler := mocks.NewMockHandler(mockService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /mocks", mockHandler.Post)

	mux.HandleFunc("GET /mocks", mockHandler.GetByParams)

	http.ListenAndServe(":9090", mux)
}
