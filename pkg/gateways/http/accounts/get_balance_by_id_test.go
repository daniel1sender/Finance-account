package accounts

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	accounts_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	accounts_storage "github.com/daniel1sender/Desafio-API/pkg/gateways/store/accounts"
)

func TestGetByID(t *testing.T) {

	t.Run("should return 200 and the account balance", func(t *testing.T) {

		storage := accounts_storage.NewStorage()
		useCase := accounts_usecase.NewUseCase(storage)
		h := NewHandler(useCase)

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, _ := useCase.Create(name, cpf, secret, balance)

		requestByID := RequestByID{account.ID}

		request, _ := json.Marshal(requestByID)

		newRequest, _ := http.NewRequest(http.MethodGet, "accounts/{id}", bytes.NewReader(request))
		newResponse := httptest.NewRecorder()

		h.GetBalanceByID(newResponse, newRequest)

		var response ResponseByID

		_ = json.Unmarshal(newResponse.Body.Bytes(), &response)

		if newResponse.Code != http.StatusCreated {
			t.Errorf("expected '%d' but got '%d'", http.StatusCreated, newResponse.Code)
		}

		if response.Balance != account.Balance {
			t.Errorf("expected '%d' but got '%d'", account.Balance, response.Balance)
		}

	})

	t.Run("should return 400 and a error message when marshalling json passed failed", func(t *testing.T) {

		storage := accounts_storage.NewStorage()
		useCase := accounts_usecase.NewUseCase(storage)
		h := NewHandler(useCase)

		b := []byte{}
		newResquest, _ := http.NewRequest(http.MethodGet, "/accounts/{id}", bytes.NewBuffer(b))
		newResponse := httptest.NewRecorder()

		h.GetBalanceByID(newResponse, newResquest)

		var response Error
		_ = json.Unmarshal(newResponse.Body.Bytes(), &response)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		expected := "invalid request body"
		if response.Reason != expected {
			t.Errorf("expected '%s' but got '%s'", expected, response.Reason)
		}

	})

	t.Run("should return 404 and a error message when account is not found by id", func(t *testing.T) {

		storage := accounts_storage.NewStorage()
		useCase := accounts_usecase.NewUseCase(storage)
		h := NewHandler(useCase)

		RequestByID := RequestByID{""}

		request, _ := json.Marshal(RequestByID)

		newRequest, _ := http.NewRequest(http.MethodGet, "accounts/{id}", bytes.NewReader(request))
		newResponse := httptest.NewRecorder()

		h.GetBalanceByID(newResponse, newRequest)

		var response ResponseByID

		_ = json.Unmarshal(newResponse.Body.Bytes(), &response)

		if newResponse.Code != http.StatusNotFound {
			t.Errorf("expected '%d' but got '%d'", http.StatusNotFound, newResponse.Code)
		}

		if response.Balance != 0 {
			t.Errorf("expected null balance but got '%d'", response.Balance)
		}

	})

}
