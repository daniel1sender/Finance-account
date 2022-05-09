package accounts

import (
	"errors"
	"net/http"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
	"github.com/sirupsen/logrus"
)

type Account struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	Balance   int    `json:"balance"`
}

type GetAccountsResponse struct {
	List []Account `json:"list"`
}

func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {

	log := h.logger
	
	accountsList, err := h.useCase.GetAll(r.Context())
	if len(accountsList) == 0 && err != nil {
		var statusCode int
		var responseError GetAccountsResponse
		switch {
		case errors.Is(err, accounts.ErrAccountNotFound):
			statusCode = http.StatusNotFound
			responseError = GetAccountsResponse{[]Account{}}
		default:
			statusCode = http.StatusInternalServerError
			responseError = GetAccountsResponse{[]Account{}}
		}
		_ = server_http.SendResponse(w, responseError, statusCode)
		log.WithFields(logrus.Fields{
			"status_code": statusCode,
		}).WithError(err).Error("listing all accounts request failed")
		return
	}

	response := GetAccountsResponse{}
	for _, value := range accountsList {
		account := Account{value.ID, value.Name, value.CreatedAt.Format(server_http.DateLayout), value.Balance}
		response.List = append(response.List, account)
	}

	listOfAccountsResponse := GetAccountsResponse{response.List}
	_ = server_http.SendResponse(w, listOfAccountsResponse, http.StatusOK)
	log.WithFields(logrus.Fields{
		"number_of_accounts": len(accountsList),
		"status_code":    http.StatusOK,
	}).Info("accounts listed successfully")
}
