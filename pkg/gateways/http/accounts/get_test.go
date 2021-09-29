package accounts

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"

	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
	accounts_storage "github.com/daniel1sender/Desafio-API/pkg/gateways/store/accounts"
)

func TestGet(t *testing.T) {

	t.Run("should return 200 and the list of accounts", func(t *testing.T) {

		storage := accounts_storage.NewStorage()
		useCase := accounts.NewUseCase(storage)

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		_, _ = useCase.Create(name, cpf, secret, balance)

		newRequest, _ := http.NewRequest(http.MethodGet, "/accounts", nil)
		newResponse := httptest.NewRecorder()
		h := NewHandler(useCase)
		h.Get(newResponse, newRequest)

		var accountsList ResponseGet
		_ = json.Unmarshal(newResponse.Body.Bytes(), &accountsList)

		for _, value := range accountsList.List {
			if value.Name != name {
				t.Errorf("expected '%s' but got '%s'", name, value.Name)
			}
			if value.CPF != cpf {
				t.Errorf("expected '%s' but got '%s'", cpf, value.CPF)
			}
			if value.Balance != balance {
				t.Errorf("expected '%d' but got '%d'", balance, value.Balance)
			}
		}

		if newResponse.Header().Get("content-type") != server_http.ContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.ContentType, newResponse.Header().Get("content-type"))
		}

		if newResponse.Code != http.StatusOK {
			t.Errorf("expected '%d' but got '%d'", http.StatusOK, newResponse.Code)
		}

	})

	t.Run("should return 404 and an empty list of accounts when no account was created", func(t *testing.T) {

		storage := accounts_storage.NewStorage()
		useCase := accounts.NewUseCase(storage)

		newRequest, _ := http.NewRequest(http.MethodGet, "/accounts", nil)
		newResponse := httptest.NewRecorder()
		h := NewHandler(useCase)
		h.Get(newResponse, newRequest)

		var accountsList ResponseGet
		_ = json.Unmarshal(newResponse.Body.Bytes(), &accountsList)

		if newResponse.Code != http.StatusNotFound {
			t.Errorf("expected '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		if len(accountsList.List) != 0 {
			t.Errorf("expected empty list of accounts but got '%v'", accountsList.List)
		}

		if newResponse.Header().Get("content-type") != server_http.ContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.ContentType, newResponse.Header().Get("content-type"))
		}

	})

}
