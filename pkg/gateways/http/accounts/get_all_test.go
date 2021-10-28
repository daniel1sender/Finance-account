package accounts

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
	"github.com/sirupsen/logrus"
)

func TestGet(t *testing.T) {
	log := logrus.NewEntry(logrus.New())
	log.Logger.SetFormatter(&logrus.JSONFormatter{})
	t.Run("should return 200 and the list of accounts", func(t *testing.T) {

		account := entities.Account{Name: "Jonh Doe", CPF: "12345678910", Secret: "123", Balance: 0}
		useCase := accounts.UseCaseMock{List: []entities.Account{account}}
		h := NewHandler(&useCase, log)
		newRequest, _ := http.NewRequest(http.MethodGet, "/accounts", nil)
		newRequest.Header.Add("Request-Id", "request-id")
		newResponse := httptest.NewRecorder()
		h.GetAll(newResponse, newRequest)
		ExpectedCreateAt := account.CreatedAt.Format(server_http.DateLayout)

		var accountsList GetResponse
		json.Unmarshal(newResponse.Body.Bytes(), &accountsList)

		for _, value := range accountsList.List {
			if value.Name != account.Name {
				t.Errorf("expected '%s' but got '%s'", account.Name, value.Name)
			}
			if value.ID != account.ID {
				t.Errorf("expected '%s' but got '%s'", account.ID, value.ID)
			}
			if value.CreatedAt != ExpectedCreateAt {
				t.Errorf("expected '%s' but got '%s'", value.CreatedAt, account.CreatedAt)
			}
			if value.Balance != account.Balance {
				t.Errorf("expected '%d' but got '%d'", account.Balance, value.Balance)
			}
		}

		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}

		if newResponse.Code != http.StatusOK {
			t.Errorf("expected '%d' but got '%d'", http.StatusOK, newResponse.Code)
		}

	})

	t.Run("should return 400 and an error message when no request-id was found in request header", func(t *testing.T) {

		useCase := accounts.UseCaseMock{}
		newRequest, _ := http.NewRequest(http.MethodGet, "/accounts", nil)
		newResponse := httptest.NewRecorder()
		h := NewHandler(&useCase, log)
		h.GetAll(newResponse, newRequest)

		var responseReason server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		expected := "invalid request header"
		if responseReason.Reason != expected {
			t.Errorf("expected '%s' but got '%s'", expected, responseReason.Reason)
		}

		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}

	})

	t.Run("should return 200 and an empty list of accounts when no account was created", func(t *testing.T) {

		useCase := accounts.UseCaseMock{List: []entities.Account{}}
		newRequest, _ := http.NewRequest(http.MethodGet, "/accounts", nil)
		newRequest.Header.Add("Request-Id", "request-id")
		newResponse := httptest.NewRecorder()
		h := NewHandler(&useCase, log)
		h.GetAll(newResponse, newRequest)

		var accountsList GetResponse
		json.Unmarshal(newResponse.Body.Bytes(), &accountsList)

		if newResponse.Code != http.StatusConflict {
			t.Errorf("expected '%d' but got '%d'", http.StatusConflict, newResponse.Code)
		}

		if len(accountsList.List) != 0 {
			t.Errorf("expected empty list of accounts but got '%v'", accountsList.List)
		}

		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}

	})

}
