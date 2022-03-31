package login

import (
	"bytes"
	"encoding/json"
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
		if newResponse.Code != http.StatusCreated {
			t.Errorf("expected '%d' but got '%d'", http.StatusCreated, newResponse.Code)
		}
		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}
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
		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected status '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}
		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}
		expected := "invalid request body"
		if response.Reason != expected {
			t.Errorf("expected '%s' but got '%s'", expected, response.Reason)
		}
	})
	t.Run("should return 404 and an error when account is not found", func(t *testing.T) {
		useCase := login.UseCaseMock{Error: accounts.ErrAccountNotFound}
		handler := NewHandler(&useCase, log)
		requestBody := Request{"12345678910", "123"}
		request, _ := json.Marshal(requestBody)
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(request))
		newResponse := httptest.NewRecorder()
		handler.Login(newResponse, newRequest)
		var response server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &response)
		if newResponse.Code != http.StatusNotFound {
			t.Errorf("expected '%d' but got '%d'", http.StatusNotFound, newResponse.Code)
		}
		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}
		if response.Reason != accounts.ErrAccountNotFound.Error() {
			t.Errorf("expected '%s' but got '%s'", accounts.ErrAccountNotFound.Error(), response.Reason)
		}
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
		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected '%d' but got '%d'", http.StatusNotFound, newResponse.Code)
		}
		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}
		if response.Reason != login.ErrEmptySecret.Error() {
			t.Errorf("expected '%s' but got '%s'", login.ErrEmptySecret.Error(), response.Reason)
		}
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
		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected '%d' but got '%d'", http.StatusNotFound, newResponse.Code)
		}
		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}
		if response.Reason != login.ErrInvalidCPF.Error() {
			t.Errorf("expected '%s' but got '%s'", login.ErrInvalidCPF.Error(), response.Reason)
		}
	})
	t.Run("should return 400 and an error when secret informed is invalid", func(t *testing.T) {
		useCase := login.UseCaseMock{Error: login.ErrInvalidSecret}
		handler := NewHandler(&useCase, log)
		requestBody := Request{"12345678910", "123"}
		request, _ := json.Marshal(requestBody)
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(request))
		newResponse := httptest.NewRecorder()
		handler.Login(newResponse, newRequest)
		var response server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &response)
		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected '%d' but got '%d'", http.StatusNotFound, newResponse.Code)
		}
		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}
		if response.Reason != login.ErrInvalidSecret.Error() {
			t.Errorf("expected '%s' but got '%s'", login.ErrInvalidSecret.Error(), response.Reason)
		}
	})
}
