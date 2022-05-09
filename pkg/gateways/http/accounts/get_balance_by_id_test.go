package accounts

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	accounts_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
	"github.com/sirupsen/logrus"
)

func TestHandlerGetBalanceByID(t *testing.T) {
	log := logrus.NewEntry(logrus.New())
	t.Run("should return 200 and the account balance", func(t *testing.T) {

		expectedBalance := 20
		useCase := accounts_usecase.UseCaseMock{Balance: expectedBalance, Error: nil}
		h := NewHandler(&useCase, log)

		newRequest, _ := http.NewRequest(http.MethodGet, "accounts/{id}/balance", nil)
		newResponse := httptest.NewRecorder()

		h.GetBalanceByID(newResponse, newRequest)

		var response GetBalanceByIdResponse
		json.Unmarshal(newResponse.Body.Bytes(), &response)

		if newResponse.Code != http.StatusOK {
			t.Errorf("expected %d but got %d", http.StatusOK, newResponse.Code)
		}

		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected %s but got %s", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}

		if response.Balance != expectedBalance {
			t.Errorf("expected '%d' but got '%d'", expectedBalance, response.Balance)
		}

	})

	t.Run("should return 404 and an error when account is not found by id", func(t *testing.T) {

		expectedBalance := 0
		useCase := accounts_usecase.UseCaseMock{Balance: expectedBalance,
			Error: accounts_usecase.ErrAccountNotFound}
		h := NewHandler(&useCase, log)

		newRequest, _ := http.NewRequest(http.MethodGet, "accounts/{id}/balance", nil)
		newResponse := httptest.NewRecorder()

		h.GetBalanceByID(newResponse, newRequest)

		var response server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &response)

		if newResponse.Code != http.StatusNotFound {
			t.Errorf("expected '%d' but got '%d'", http.StatusNotFound, newResponse.Code)
		}

		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}

		if response.Reason != accounts_usecase.ErrAccountNotFound.Error() {
			t.Errorf("expected '%s' but got '%s'", accounts_usecase.ErrAccountNotFound.Error(), response.Reason)
		}

	})

}
