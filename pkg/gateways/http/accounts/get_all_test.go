package accounts

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
)

func TestHandlerGetAll(t *testing.T) {
	log := logrus.NewEntry(logrus.New())
	t.Run("should return 200 and the list of accounts", func(t *testing.T) {

		account := entities.Account{Name: "Jonh Doe", CPF: "12345678910", Secret: "123", Balance: 0}
		useCase := accounts.UseCaseMock{List: []entities.Account{account}}

		h := NewHandler(&useCase, log)

		newRequest, _ := http.NewRequest(http.MethodGet, "/accounts", nil)
		newResponse := httptest.NewRecorder()
		h.GetAll(newResponse, newRequest)

		expectedCreateAt := account.CreatedAt.Format(server_http.DateLayout)

		var accountsList GetAccountsResponse
		json.Unmarshal(newResponse.Body.Bytes(), &accountsList)

		for _, value := range accountsList.List {
			assert.Equal(t, value.Name, account.Name)
			assert.Equal(t, value.ID, account.ID)
			assert.Equal(t, value.CreatedAt, expectedCreateAt)
			assert.Equal(t, value.Balance, account.Balance)
		}
		assert.Equal(t, newResponse.Header().Get("content-type"), server_http.JSONContentType)
		assert.Equal(t, newResponse.Code, http.StatusOK)
	})

	t.Run("should return 404 and an empty list of accounts when no account was created", func(t *testing.T) {

		useCase := accounts.UseCaseMock{List: []entities.Account{}, Error: accounts.ErrAccountNotFound}
		newRequest, _ := http.NewRequest(http.MethodGet, "/accounts", nil)
		newResponse := httptest.NewRecorder()

		h := NewHandler(&useCase, log)

		h.GetAll(newResponse, newRequest)

		var accountsList GetAccountsResponse
		json.Unmarshal(newResponse.Body.Bytes(), &accountsList)

		assert.Equal(t, newResponse.Code, http.StatusNotFound)
		assert.Empty(t, accountsList.List)
		assert.Equal(t, newResponse.Header().Get("content-type"), server_http.JSONContentType)
	})

	t.Run("should return 500 and an empty list of accounts when some error with database occur", func(t *testing.T) {
		useCase := accounts.UseCaseMock{List: []entities.Account{}, Error: accounts.ErrEmptyList}
		newRequest, _ := http.NewRequest(http.MethodGet, "/accounts", nil)
		newResponse := httptest.NewRecorder()

		h := NewHandler(&useCase, log)

		h.GetAll(newResponse, newRequest)

		var accountsList GetAccountsResponse
		json.Unmarshal(newResponse.Body.Bytes(), &accountsList)

		assert.Equal(t, newResponse.Code, http.StatusInternalServerError)
		assert.Empty(t, accountsList.List)
		assert.Equal(t, newResponse.Header().Get("content-type"), server_http.JSONContentType)
	})

}
