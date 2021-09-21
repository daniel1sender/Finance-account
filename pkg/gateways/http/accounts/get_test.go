package accounts

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"

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

		newRequest, _ := http.NewRequest(http.MethodGet, "/anyroute", nil)
		newResponse := httptest.NewRecorder()
		h := NewHandler(useCase)
		h.Get(newResponse, newRequest)

		var accountsList Response
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

	})

}
