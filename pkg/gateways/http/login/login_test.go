package login

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/login"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestHandlerLogin(t *testing.T) {
	log := logrus.NewEntry(logrus.New())

	t.Run("should return 201 and the token created", func(t *testing.T) {
		useCase := login.UseCaseMock{
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJjM2Q3OGU2My1mYWQ5LTQ1ZTgtODI0OC1iOGNjMDc4ZDFiZGYiLCJleHAiOjE2NDg2Nzc5MDYsImlhdCI6MTY0ODY3NzYwNiwianRpIjoiZDNhNDYxMDctMDQzNi00ZmViLTk4YmItZGEzMTQzZmFiYWQyIn0.cN2-NOvwMwvEIKxBBVV-toJkmohkRDwppzCQs7XOqpA",
			Error: nil,
		}
		handler := NewHandler(&useCase, log)
		requestBody := Request{"12345678910", "123"}
		request, _ := json.Marshal(requestBody)
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(request))
		newResponse := httptest.NewRecorder()
		handler.Login(newResponse, newRequest)
		var response Response
		json.Unmarshal(newResponse.Body.Bytes(), &response)

		assert.Equal(t, newResponse.Code, http.StatusCreated)
		assert.Equal(t, newResponse.Header().Get("content-type"), server_http.JSONContentType)
		assert.NotEmpty(t, response.Token)
		assert.Equal(t, response.Token, useCase.Token)
	})

	t.Run("should return 400 and an error when it failed to decode the request successfully", func(t *testing.T) {

		useCase := login.UseCaseMock{}
		h := NewHandler(&useCase, log)
		b := []byte{}
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(b))
		newResponse := httptest.NewRecorder()
		h.Login(newResponse, newRequest)
		var response server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &response)

		assert.Equal(t, newResponse.Code, http.StatusBadRequest)
		assert.Equal(t, newResponse.Header().Get("content-type"), server_http.JSONContentType)
		expected := "invalid request body"
		assert.Equal(t, response.Reason, expected)
	})

	t.Run("should return 404 and an error when the account is not found", func(t *testing.T) {
		useCase := login.UseCaseMock{Error: accounts.ErrAccountNotFound}
		handler := NewHandler(&useCase, log)
		requestBody := Request{"12345678910", "123"}
		request, _ := json.Marshal(requestBody)
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(request))
		newResponse := httptest.NewRecorder()
		handler.Login(newResponse, newRequest)
		var response server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &response)

		assert.Equal(t, newResponse.Code, http.StatusForbidden)
		assert.Equal(t, newResponse.Header().Get("content-type"), server_http.JSONContentType)
		assert.Equal(t, response.Reason, login.ErrInvalidCredentials.Error())
	})

	t.Run("should return 400 and an error when an empty secret is informed", func(t *testing.T) {
		useCase := login.UseCaseMock{Error: login.ErrEmptySecret}
		handler := NewHandler(&useCase, log)
		requestBody := Request{"12345678910", "123"}
		request, _ := json.Marshal(requestBody)
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(request))
		newResponse := httptest.NewRecorder()
		handler.Login(newResponse, newRequest)
		var response server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &response)

		assert.Equal(t, newResponse.Code, http.StatusBadRequest)
		assert.Equal(t, newResponse.Header().Get("content-type"), server_http.JSONContentType)
		assert.Equal(t, response.Reason, login.ErrEmptySecret.Error())
	})

	t.Run("should return 400 and an error when the cpf informed doesn't have eleven digits", func(t *testing.T) {
		useCase := login.UseCaseMock{Error: login.ErrInvalidCPF}
		handler := NewHandler(&useCase, log)
		requestBody := Request{"12345678910", "123"}
		request, _ := json.Marshal(requestBody)
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(request))
		newResponse := httptest.NewRecorder()
		handler.Login(newResponse, newRequest)
		var response server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &response)

		assert.Equal(t, newResponse.Code, http.StatusBadRequest)
		assert.Equal(t, newResponse.Header().Get("content-type"), server_http.JSONContentType)
		assert.Equal(t, response.Reason, login.ErrInvalidCPF.Error())
	})

	t.Run("should return 400 and an error when the secret informed is invalid", func(t *testing.T) {
		useCase := login.UseCaseMock{Error: login.ErrInvalidSecret}
		handler := NewHandler(&useCase, log)
		requestBody := Request{"12345678910", "123"}
		request, _ := json.Marshal(requestBody)
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(request))
		newResponse := httptest.NewRecorder()
		handler.Login(newResponse, newRequest)
		var response server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &response)

		assert.Equal(t, newResponse.Code, http.StatusForbidden)
		assert.Equal(t, newResponse.Header().Get("content-type"), server_http.JSONContentType)
		assert.Equal(t, response.Reason, login.ErrInvalidCredentials.Error())
	})

	t.Run("should return 500 and an error when an unexpected error occourred", func(t *testing.T) {
		unexpectedError := errors.New("unexpected error")
		useCase := login.UseCaseMock{Error: unexpectedError}
		handler := NewHandler(&useCase, log)
		requestBody := Request{"12345678910", "123"}
		request, _ := json.Marshal(requestBody)
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(request))
		newResponse := httptest.NewRecorder()
		handler.Login(newResponse, newRequest)
		var response server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &response)

		assert.Equal(t, newResponse.Code, http.StatusInternalServerError)
		assert.Equal(t, newResponse.Header().Get("content-type"), server_http.JSONContentType)
		expectedReason := "internal server error"
		assert.Equal(t, response.Reason, expectedReason)
	})
}
