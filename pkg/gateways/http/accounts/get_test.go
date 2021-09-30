package accounts

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"

	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
)

func TestGet(t *testing.T) {

	t.Run("should return 200 and the list of accounts", func(t *testing.T) {

		account := entities.Account{Name: "Jonh Doe", CPF: "12345678910", Secret: "123", Balance: 0}

		useCase := accounts.UseCaseMock{Balance: 0, Error: nil, Account: account}

		h := NewHandler(&useCase)

		useCase.Create(account.Name, account.CPF, account.Secret, account.Balance)

		newRequest, _ := http.NewRequest(http.MethodGet, "/accounts", nil)
		newResponse := httptest.NewRecorder()
		h.Get(newResponse, newRequest)

		var accountsList GetResponse
		_ = json.Unmarshal(newResponse.Body.Bytes(), &accountsList)

		for _, value := range accountsList.List {
			if value.Name != account.Name {
				t.Errorf("expected '%s' but got '%s'", account.Name, value.Name)
			}
			if value.CPF != account.CPF {
				t.Errorf("expected '%s' but got '%s'", account.CPF, value.CPF)
			}
			if value.Balance != account.Balance {
				t.Errorf("expected '%d' but got '%d'", account.Balance, value.Balance)
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

		useCase := accounts.UseCaseMock{Balance: 0, Error: nil}

		newRequest, _ := http.NewRequest(http.MethodGet, "/accounts", nil)
		newResponse := httptest.NewRecorder()
		h := NewHandler(&useCase)

		h.Get(newResponse, newRequest)

		var accountsList GetResponse
		_ = json.Unmarshal(newResponse.Body.Bytes(), &accountsList)

		if newResponse.Code != http.StatusNotFound {
			t.Errorf("expected '%d' but got '%d'", http.StatusNotFound, newResponse.Code)
		}

		if len(accountsList.List) != 0 {
			t.Errorf("expected empty list of accounts but got '%v'", accountsList.List)
		}

		if newResponse.Header().Get("content-type") != server_http.ContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.ContentType, newResponse.Header().Get("content-type"))
		}

	})

}
