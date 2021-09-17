package accounts

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	accounts_storage "github.com/daniel1sender/Desafio-API/pkg/gateways/store/accounts"
)

func TestCreate(t *testing.T) {

	type createRequest struct {
		Name    string
		CPF     string
		Secret  string
		Balance int
	}

	t.Run("should return 200 and null error when the type informed is json", func(t *testing.T) {

		request := createRequest{"Jonh Doe", "12345678910", "123", 0}

		requestBody, _ := json.Marshal(request)

		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(requestBody))
		newResponse := httptest.NewRecorder()

		storage := accounts_storage.NewStorage()
		useCase := accounts.NewUseCase(storage)
		h := NewHandler(useCase)

		h.Create(newResponse, newRequest)

		var responseValidation CreateResponse
		_ = json.Unmarshal(newResponse.Body.Bytes(), &responseValidation)

		if newResponse.Code != http.StatusCreated {
			t.Errorf("expected '%d' but got '%d'", http.StatusCreated, newResponse.Code)
		}

		if responseValidation.Name != request.Name {
			t.Errorf("expected '%s' but got '%s'", request.Name, responseValidation.Name)
		}

		if responseValidation.CPF != request.CPF {
			t.Errorf("expected '%s' but got '%s'", request.CPF, responseValidation.CPF)
		}

		if responseValidation.Balance != request.Balance {
			t.Errorf("expected '%d' but got '%d'", 0, responseValidation.Balance)
		}

		if responseValidation.CreatedAt.IsZero() {
			t.Errorf("expected nonzero time but got '%s'", responseValidation.CreatedAt)
		}

	})

	t.Run("should return 400 and a error message when the type informed it is not a json", func(t *testing.T) {

		b := []byte{}
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(b))
		newResponse := httptest.NewRecorder()

		storage := accounts_storage.NewStorage()
		useCase := accounts.NewUseCase(storage)
		h := NewHandler(useCase)

		h.Create(newResponse, newRequest)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected status '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		var responseReason Error
		_ = json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		expected := "invalid request body"
		if responseReason.Reason != expected {
			t.Errorf("expected '%s' but got '%s'", expected, responseReason.Reason)
		}

	})

	t.Run("should return 400 and a message error when an empty name is informed", func(t *testing.T) {

		request := createRequest{"", "12345678910", "123", 0}
		requestBody, _ := json.Marshal(request)
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(requestBody))
		newResponse := httptest.NewRecorder()

		storage := accounts_storage.NewStorage()
		useCase := accounts.NewUseCase(storage)
		h := NewHandler(useCase)

		h.Create(newResponse, newRequest)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected status '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		var responseReason Error
		_ = json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		expected := "error while creating an account"
		if responseReason.Reason != expected {
			t.Errorf("expected '%s' but got '%s'", expected, responseReason.Reason)
		}

	})

	t.Run("should return 400 and a message error when the cpf informed doesn't have eleven digits", func(t *testing.T) {

		request := createRequest{"Jonh Doe", "1234567891", "123", 0}
		requestBody, _ := json.Marshal(request)
		newRequest, _ := http.NewRequest("POST", "/anyroute", bytes.NewReader(requestBody))
		newResponse := httptest.NewRecorder()

		storage := accounts_storage.NewStorage()
		useCase := accounts.NewUseCase(storage)
		h := NewHandler(useCase)

		h.Create(newResponse, newRequest)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		var responseReason Error
		_ = json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		expected := "error while creating an account"
		if responseReason.Reason != expected {
			t.Errorf("expected '%s' but got '%s'", expected, responseReason.Reason)
		}
	})

	t.Run("should return a 400 and a message error when cpf informed already exist", func(t *testing.T) {

		request := createRequest{"Jonh Doe", "12345678910", "123", 0}
		requestBody, _ := json.Marshal(request)
		newRequest, _ := http.NewRequest("POST", "anyroute", bytes.NewReader(requestBody))
		newResponse := httptest.NewRecorder()

		storage := accounts_storage.NewStorage()
		useCase := accounts.NewUseCase(storage)
		h := NewHandler(useCase)

		h.Create(newResponse, newRequest)

		request = createRequest{"Jonh Doe", "12345678910", "123", 0}
		requestBody, _ = json.Marshal(request)
		newRequest, _ = http.NewRequest("POST", "anyroute", bytes.NewReader(requestBody))
		newResponse = httptest.NewRecorder()

		h.Create(newResponse, newRequest)

		var responseReason Error
		_ = json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		if responseReason.Reason != accounts.ErrExistingCPF.Error() {
			t.Errorf("expected '%s' but got '%s'", accounts.ErrExistingCPF.Error(), responseReason.Reason)
		}

	})

	t.Run("should return a 400 and a message error when a blanc secret is informed", func(t *testing.T) {

		request := createRequest{"Jonh Doe", "12345678910", "", 0}
		requestBody, _ := json.Marshal(request)
		newRequest, _ := http.NewRequest("POST", "anyroute", bytes.NewReader(requestBody))
		newResponse := httptest.NewRecorder()

		storage := accounts_storage.NewStorage()
		useCase := accounts.NewUseCase(storage)
		h := NewHandler(useCase)

		h.Create(newResponse, newRequest)

		var responseReason Error
		_ = json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		expected := "error while creating an account"
		if responseReason.Reason != expected {
			t.Errorf("expected '%s' but got '%s'", expected, responseReason.Reason)
		}

	})

}
