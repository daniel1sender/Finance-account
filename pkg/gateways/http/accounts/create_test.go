package accounts

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
	accounts_storage "github.com/daniel1sender/Desafio-API/pkg/gateways/store/accounts"
)

func TestCreate(t *testing.T) {

	t.Run("should return 200 and null error when the type informed is json", func(t *testing.T) {

		storage := accounts_storage.NewStorage()
		useCase := accounts.NewUseCase(storage)
		h := NewHandler(useCase)

		requestCreate := RequestCreate{"Jonh Doe", "12345678910", "123", 0}

		request, _ := json.Marshal(requestCreate)

		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(request))
		newResponse := httptest.NewRecorder()

		h.Create(newResponse, newRequest)

		var response ResponseCreate
		_ = json.Unmarshal(newResponse.Body.Bytes(), &response)

		if newResponse.Code != http.StatusCreated {
			t.Errorf("expected '%d' but got '%d'", http.StatusCreated, newResponse.Code)
		}

		if newResponse.Header().Get("content-type") != server_http.ContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.ContentType, newResponse.Header().Get("content-type"))
		}

		if response.Name != requestCreate.Name {
			t.Errorf("expected '%s' but got '%s'", requestCreate.Name, response.Name)
		}

		if response.CPF != requestCreate.CPF {
			t.Errorf("expected '%s' but got '%s'", requestCreate.CPF, response.CPF)
		}

		if response.Balance != requestCreate.Balance {
			t.Errorf("expected '%d' but got '%d'", 0, response.Balance)
		}

		if response.CreatedAt.IsZero() {
			t.Errorf("expected nonzero time but got '%s'", response.CreatedAt)
		}

	})

	t.Run("should return 400 and a error message when the type informed it is not a json", func(t *testing.T) {

		storage := accounts_storage.NewStorage()
		useCase := accounts.NewUseCase(storage)
		h := NewHandler(useCase)

		b := []byte{}
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(b))
		newResponse := httptest.NewRecorder()

		h.Create(newResponse, newRequest)

		var responseReason Error
		_ = json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected status '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		if newResponse.Header().Get("content-type") != server_http.ContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.ContentType, newResponse.Header().Get("content-type"))
		}

		expected := "invalid request body"
		if responseReason.Reason != expected {
			t.Errorf("expected '%s' but got '%s'", expected, responseReason.Reason)
		}

	})

	t.Run("should return 400 and a message error when an empty name is informed", func(t *testing.T) {

		storage := accounts_storage.NewStorage()
		useCase := accounts.NewUseCase(storage)
		h := NewHandler(useCase)

		requestCreate := RequestCreate{"", "12345678910", "123", 0}
		request, _ := json.Marshal(requestCreate)
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(request))
		newResponse := httptest.NewRecorder()

		h.Create(newResponse, newRequest)

		var responseReason Error
		_ = json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected status '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		if newResponse.Header().Get("content-type") != server_http.ContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.ContentType, newResponse.Header().Get("content-type"))
		}

		if responseReason.Reason != entities.ErrInvalidName.Error() {
			t.Errorf("expected '%s' but got '%s'", entities.ErrInvalidName, responseReason.Reason)
		}
	})

	t.Run("should return 400 and a message error when the cpf informed doesn't have eleven digits", func(t *testing.T) {

		storage := accounts_storage.NewStorage()
		useCase := accounts.NewUseCase(storage)
		h := NewHandler(useCase)

		requestCreate := RequestCreate{"Jonh Doe", "1234567891", "123", 0}
		requestBody, _ := json.Marshal(requestCreate)
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(requestBody))
		newResponse := httptest.NewRecorder()

		h.Create(newResponse, newRequest)

		var responseReason Error
		_ = json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		if newResponse.Header().Get("content-type") != server_http.ContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.ContentType, newResponse.Header().Get("content-type"))
		}

		if responseReason.Reason != entities.ErrInvalidCPF.Error() {
			t.Errorf("expected '%s' but got '%s'", entities.ErrInvalidCPF.Error(), responseReason.Reason)
		}

	})

	t.Run("should return a 400 and a message error when cpf informed already exist", func(t *testing.T) {

		storage := accounts_storage.NewStorage()
		useCase := accounts.NewUseCase(storage)
		h := NewHandler(useCase)

		requestCreate := RequestCreate{"Jonh Doe", "12345678910", "123", 0}
		requestBody, _ := json.Marshal(requestCreate)
		newRequest, _ := http.NewRequest("POST", "anyroute", bytes.NewReader(requestBody))
		newResponse := httptest.NewRecorder()

		h.Create(newResponse, newRequest)

		requestCreate = RequestCreate{"Jonh Doe", "12345678910", "123", 0}
		requestBody, _ = json.Marshal(requestCreate)
		newRequest, _ = http.NewRequest("POST", "anyroute", bytes.NewReader(requestBody))
		newResponse = httptest.NewRecorder()

		h.Create(newResponse, newRequest)

		var responseReason Error
		_ = json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		if newResponse.Header().Get("content-type") != server_http.ContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.ContentType, newResponse.Header().Get("content-type"))
		}

		if responseReason.Reason != accounts.ErrExistingCPF.Error() {
			t.Errorf("expected '%s' but got '%s'", accounts.ErrExistingCPF.Error(), responseReason.Reason)
		}

	})

	t.Run("should return 400 and a message error when a blanc secret is informed", func(t *testing.T) {

		storage := accounts_storage.NewStorage()
		useCase := accounts.NewUseCase(storage)
		h := NewHandler(useCase)

		requestCreate := RequestCreate{"Jonh Doe", "12345678910", "", 0}
		requestBody, _ := json.Marshal(requestCreate)
		newRequest, _ := http.NewRequest("POST", "anyroute", bytes.NewReader(requestBody))
		newResponse := httptest.NewRecorder()

		h.Create(newResponse, newRequest)

		var responseReason Error
		_ = json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		if newResponse.Header().Get("content-type") != server_http.ContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.ContentType, newResponse.Header().Get("content-type"))
		}

		if responseReason.Reason != entities.ErrBlankSecret.Error() {
			t.Errorf("expected '%s' but got '%s'", entities.ErrBlankSecret.Error(), responseReason.Reason)
		}

	})

	t.Run("should return 400 and a message error when balance informed is less than zero", func(t *testing.T) {

		storage := accounts_storage.NewStorage()
		useCase := accounts.NewUseCase(storage)
		h := NewHandler(useCase)

		requestCreate := RequestCreate{"Jonh Doe", "12345678910", "123", -10}
		requestBody, _ := json.Marshal(requestCreate)
		newRequest, _ := http.NewRequest("POST", "anyroute", bytes.NewReader(requestBody))
		newResponse := httptest.NewRecorder()

		h.Create(newResponse, newRequest)

		var responseReason Error
		_ = json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		if newResponse.Header().Get("content-type") != server_http.ContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.ContentType, newResponse.Header().Get("content-type"))
		}

		if responseReason.Reason != entities.ErrBalanceLessZero.Error() {
			t.Errorf("expected '%s' but got '%s'", entities.ErrBalanceLessZero.Error(), responseReason.Reason)
		}

	})

}
