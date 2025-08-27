package main

import (
	"log"
	"net/http"

	"github.com/eduufreire/poc-mock-services/internal/aws"
	"github.com/eduufreire/poc-mock-services/internal/aws/dynamo"
	"github.com/eduufreire/poc-mock-services/internal/cache"
	"github.com/eduufreire/poc-mock-services/internal/mocks"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	clientDynamo := dynamo.NewDynamoService(aws.Session())
	clientCache := cache.NewClientRedis()

	mockService := mocks.NewMockService(clientDynamo, clientCache)
	mockHandler := mocks.NewMockHandler(mockService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /mocks", mockHandler.Post)

	mux.HandleFunc("GET /mocks", mockHandler.GetByParams)

	http.ListenAndServe(":9090", mux)
}
