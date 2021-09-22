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

		newRequest, _ := http.NewRequest("GET", "accounts/{id}", bytes.NewReader(request))
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

}
