package accounts

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
)

func TestCreate(t *testing.T) {
	log := server_http.NewLogger()
	t.Run("should return 201 and a account when it's been sucessfully created", func(t *testing.T) {

		account := entities.Account{Name: "Jonh Doe", CPF: "12345678910", Secret: "123", Balance: 0, CreatedAt: time.Now()}
		useCase := accounts.UseCaseMock{Account: account}
		h := NewHandler(&useCase, log)
		createRequest := CreateRequest{account.Name, account.CPF, account.Secret, account.Balance}
		request, _ := json.Marshal(createRequest)
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(request))
		newRequest.Header.Add("Request-Id", server_http.KeyHeader)
		newResponse := httptest.NewRecorder()

		h.Create(newResponse, newRequest)

		ExpectedCreateAt := account.CreatedAt.Format(server_http.DateLayout)

		var response CreateResponse
		json.Unmarshal(newResponse.Body.Bytes(), &response)

		if newResponse.Code != http.StatusCreated {
			t.Errorf("expected '%d' but got '%d'", http.StatusCreated, newResponse.Code)
		}

		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}

		if response.Name != createRequest.Name {
			t.Errorf("expected '%s' but got '%s'", createRequest.Name, response.Name)
		}

		if response.CPF != createRequest.CPF {
			t.Errorf("expected '%s' but got '%s'", createRequest.CPF, response.CPF)
		}

		if response.Balance != createRequest.Balance {
			t.Errorf("expected '%d' but got '%d'", 0, response.Balance)
		}

		if response.CreatedAt != ExpectedCreateAt {
			t.Errorf("expected '%s' but got '%s'", ExpectedCreateAt, response.CreatedAt)
		}
	})

	t.Run("should return 400 and an error message when no request-id was found in request header", func(t *testing.T) {

		useCase := accounts.UseCaseMock{}
		h := NewHandler(&useCase, log)
		createRequest := CreateRequest{}
		request, _ := json.Marshal(createRequest)
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(request))
		newResponse := httptest.NewRecorder()
		h.Create(newResponse, newRequest)

		var responseReason server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected status '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}

		expected := "invalid request header"
		if responseReason.Reason != expected {
			t.Errorf("expected '%s' but got '%s'", expected, responseReason.Reason)
		}

	})

	t.Run("should return 400 and a error message when it failed to decode the request successfully", func(t *testing.T) {

		useCase := accounts.UseCaseMock{}
		h := NewHandler(&useCase, log)
		b := []byte{}
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(b))
		newRequest.Header.Add("Request-Id", server_http.KeyHeader)
		newResponse := httptest.NewRecorder()
		h.Create(newResponse, newRequest)

		var responseReason server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected status '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}

		expected := "invalid request body"
		if responseReason.Reason != expected {
			t.Errorf("expected '%s' but got '%s'", expected, responseReason.Reason)
		}

	})

	t.Run("should return 400 and a message error when an empty name is informed", func(t *testing.T) {

		useCase := accounts.UseCaseMock{Error: entities.ErrInvalidName}
		h := NewHandler(&useCase, log)
		createRequest := CreateRequest{}
		request, _ := json.Marshal(createRequest)
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(request))
		newRequest.Header.Add("Request-Id", server_http.KeyHeader)
		newResponse := httptest.NewRecorder()

		h.Create(newResponse, newRequest)

		var responseReason server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected status '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}

		if responseReason.Reason != entities.ErrInvalidName.Error() {
			t.Errorf("expected '%s' but got '%s'", entities.ErrInvalidName, responseReason.Reason)
		}
	})

	t.Run("should return 400 and a message error when the cpf informed doesn't have eleven digits", func(t *testing.T) {

		useCase := accounts.UseCaseMock{Error: entities.ErrInvalidCPF}
		h := NewHandler(&useCase, log)
		createRequest := CreateRequest{}
		requestBody, _ := json.Marshal(createRequest)
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(requestBody))
		newRequest.Header.Add("Request-Id", server_http.KeyHeader)
		newResponse := httptest.NewRecorder()

		h.Create(newResponse, newRequest)

		var responseReason server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}

		if responseReason.Reason != entities.ErrInvalidCPF.Error() {
			t.Errorf("expected '%s' but got '%s'", entities.ErrInvalidCPF.Error(), responseReason.Reason)
		}

	})

	t.Run("should return a 409 and a message error when cpf informed already exist", func(t *testing.T) {

		useCase := accounts.UseCaseMock{Error: accounts.ErrExistingCPF}
		h := NewHandler(&useCase, log)
		createRequest := CreateRequest{}
		requestBody, _ := json.Marshal(createRequest)
		newRequest, _ := http.NewRequest("POST", "anyroute", bytes.NewReader(requestBody))
		newRequest.Header.Add("Request-Id", server_http.KeyHeader)
		newResponse := httptest.NewRecorder()

		h.Create(newResponse, newRequest)

		var responseReason server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		if newResponse.Code != http.StatusConflict {
			t.Errorf("expected '%d' but got '%d'", http.StatusConflict, newResponse.Code)
		}

		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}

		if responseReason.Reason != accounts.ErrExistingCPF.Error() {
			t.Errorf("expected '%s' but got '%s'", accounts.ErrExistingCPF.Error(), responseReason.Reason)
		}

	})

	t.Run("should return 400 and a message error when a blanc secret is informed", func(t *testing.T) {

		useCase := accounts.UseCaseMock{Error: entities.ErrEmptySecret}
		h := NewHandler(&useCase, log)
		createRequest := CreateRequest{}
		requestBody, _ := json.Marshal(createRequest)
		newRequest, _ := http.NewRequest("POST", "anyroute", bytes.NewReader(requestBody))
		newRequest.Header.Add("Request-Id", server_http.KeyHeader)
		newResponse := httptest.NewRecorder()

		h.Create(newResponse, newRequest)

		var responseReason server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}

		if responseReason.Reason != entities.ErrEmptySecret.Error() {
			t.Errorf("expected '%s' but got '%s'", entities.ErrEmptySecret.Error(), responseReason.Reason)
		}

	})

	t.Run("should return 400 and a message error when balance informed is less than zero", func(t *testing.T) {

		useCase := accounts.UseCaseMock{Error: entities.ErrNegativeBalance}
		h := NewHandler(&useCase, log)
		createRequest := CreateRequest{}
		requestBody, _ := json.Marshal(createRequest)
		newRequest, _ := http.NewRequest("POST", "anyroute", bytes.NewReader(requestBody))
		newRequest.Header.Add("Request-Id", server_http.KeyHeader)
		newResponse := httptest.NewRecorder()

		h.Create(newResponse, newRequest)

		var responseReason server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}

		if responseReason.Reason != entities.ErrNegativeBalance.Error() {
			t.Errorf("expected '%s' but got '%s'", entities.ErrNegativeBalance.Error(), responseReason.Reason)
		}

	})

}
