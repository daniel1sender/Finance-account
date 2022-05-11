package accounts

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/daniel1sender/Desafio-API/pkg/domain"
	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
	"github.com/sirupsen/logrus"
	"gotest.tools/assert"
)

func TestHandlerCreate(t *testing.T) {
	log := logrus.NewEntry(logrus.New())
	t.Run("should return 201 and an account when it's been sucessfully created", func(t *testing.T) {

		account := entities.Account{Name: "Jonh Doe", CPF: "12345678910", Secret: "123", Balance: 0, CreatedAt: time.Now()}
		useCase := accounts.UseCaseMock{Account: account}
		h := NewHandler(&useCase, log)

		createRequest := CreateAccountRequest{account.Name, account.CPF, account.Secret, account.Balance}
		request, _ := json.Marshal(createRequest)
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(request))
		newResponse := httptest.NewRecorder()

		h.Create(newResponse, newRequest)

		expectedCreateAt := account.CreatedAt.Format(server_http.DateLayout)

		var response CreateAccountResponse
		json.Unmarshal(newResponse.Body.Bytes(), &response)

		assert.Equal(t, newResponse.Code, http.StatusCreated)
		assert.Equal(t, newResponse.Header().Get("content-type"), server_http.JSONContentType)
		assert.Equal(t, response.Name, createRequest.Name)
		assert.Equal(t, response.CPF, createRequest.CPF)
		assert.Equal(t, response.Balance, createRequest.Balance)
		assert.Equal(t, response.CreatedAt, expectedCreateAt)
	})

	t.Run("should return 400 and an error when it failed to decode the request successfully", func(t *testing.T) {

		useCase := accounts.UseCaseMock{}
		h := NewHandler(&useCase, log)
		expected := "invalid request body"
		b := []byte{}
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(b))
		newResponse := httptest.NewRecorder()

		h.Create(newResponse, newRequest)

		var responseReason server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		assert.Equal(t, newResponse.Code, http.StatusBadRequest)
		assert.Equal(t, newResponse.Header().Get("content-type"), server_http.JSONContentType)
		assert.Equal(t, responseReason.Reason, expected)
	})

	t.Run("should return 400 and an error when an empty name is informed", func(t *testing.T) {

		useCase := accounts.UseCaseMock{Error: entities.ErrInvalidName}
		h := NewHandler(&useCase, log)

		createRequest := CreateAccountRequest{}
		request, _ := json.Marshal(createRequest)
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(request))
		newResponse := httptest.NewRecorder()

		h.Create(newResponse, newRequest)

		var responseReason server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		assert.Equal(t, newResponse.Code, http.StatusBadRequest)
		assert.Equal(t, newResponse.Header().Get("content-type"), server_http.JSONContentType)
		assert.Equal(t, responseReason.Reason, entities.ErrInvalidName.Error())
	})

	t.Run("should return 400 and an error when the cpf informed doesn't have eleven digits", func(t *testing.T) {

		useCase := accounts.UseCaseMock{Error: domain.ErrInvalidCPF}
		h := NewHandler(&useCase, log)

		createRequest := CreateAccountRequest{}
		requestBody, _ := json.Marshal(createRequest)
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(requestBody))
		newResponse := httptest.NewRecorder()

		h.Create(newResponse, newRequest)

		var responseReason server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		assert.Equal(t, newResponse.Code, http.StatusBadRequest)
		assert.Equal(t, newResponse.Header().Get("content-type"), server_http.JSONContentType)
		assert.Equal(t, responseReason.Reason, domain.ErrInvalidCPF.Error())
	})

	t.Run("should return 409 and an error when cpf informed already exist", func(t *testing.T) {

		useCase := accounts.UseCaseMock{Error: accounts.ErrExistingCPF}
		h := NewHandler(&useCase, log)

		createRequest := CreateAccountRequest{}
		requestBody, _ := json.Marshal(createRequest)
		newRequest, _ := http.NewRequest("POST", "anyroute", bytes.NewReader(requestBody))
		newResponse := httptest.NewRecorder()

		h.Create(newResponse, newRequest)

		var responseReason server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		assert.Equal(t, newResponse.Code, http.StatusConflict)
		assert.Equal(t, newResponse.Header().Get("content-type"), server_http.JSONContentType)
		assert.Equal(t, responseReason.Reason, accounts.ErrExistingCPF.Error())
	})

	t.Run("should return 400 and an error when an empty secret is informed", func(t *testing.T) {

		useCase := accounts.UseCaseMock{Error: domain.ErrEmptySecret}
		h := NewHandler(&useCase, log)

		createRequest := CreateAccountRequest{}
		requestBody, _ := json.Marshal(createRequest)
		newRequest, _ := http.NewRequest("POST", "anyroute", bytes.NewReader(requestBody))
		newResponse := httptest.NewRecorder()

		h.Create(newResponse, newRequest)

		var responseReason server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		assert.Equal(t, newResponse.Code, http.StatusBadRequest)
		assert.Equal(t, newResponse.Header().Get("content-type"), server_http.JSONContentType)
		assert.Equal(t, responseReason.Reason, domain.ErrEmptySecret.Error())
	})

	t.Run("should return 400 and an error when balance informed is less than zero", func(t *testing.T) {

		useCase := accounts.UseCaseMock{Error: entities.ErrNegativeBalance}
		h := NewHandler(&useCase, log)

		createRequest := CreateAccountRequest{}
		requestBody, _ := json.Marshal(createRequest)
		newRequest, _ := http.NewRequest("POST", "anyroute", bytes.NewReader(requestBody))
		newResponse := httptest.NewRecorder()

		h.Create(newResponse, newRequest)

		var responseReason server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		assert.Equal(t, newResponse.Code, http.StatusBadRequest)
		assert.Equal(t, newResponse.Header().Get("content-type"), server_http.JSONContentType)
		assert.Equal(t, responseReason.Reason, entities.ErrNegativeBalance.Error())
	})

}
