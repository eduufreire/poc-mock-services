package mocks

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func NewMockHandler(service *mockService) *mockHandler {
	return &mockHandler{
		service: service,
	}
}

func (m *mockHandler) Post(w http.ResponseWriter, r *http.Request) {
	body := Request{}
	json.NewDecoder(r.Body).Decode(&body)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	m.service.SaveMock(ctx, body)
	w.WriteHeader(201)
}

func (m *mockHandler) GetByParams(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()

	if queryParams.Get("service") == "" {
		w.WriteHeader(400)
		w.Write([]byte("param `service` required"))
		return
	}

	if queryParams.Get("statusCode") == "" {
		w.WriteHeader(400)
		w.Write([]byte("param `statusCode` required"))
		return
	}

	if queryParams.Get("endpoint") == "" {
		w.WriteHeader(400)
		w.Write([]byte("param `endpoint` required"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	statusCode, err := strconv.Atoi(queryParams.Get("statusCode"))
	if err != nil {
		log.Fatal(err)
		w.Write([]byte(err.Error()))
	}

	params := Params{
		Service:    queryParams.Get("service"),
		Endpoint:   queryParams.Get("endpoint"),
		StatusCode: statusCode,
	}
	mock := m.service.GetMockByParams(ctx, params)

	response, err := json.Marshal(mock.Payload)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(mock.StatusCode)
	w.Write(response)
}
