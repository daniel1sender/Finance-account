package accounts

import (
	"errors"
	"net/http"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type GetBalanceByIdResponse struct {
	Balance int `json:"balance"`
}

func (h Handler) GetBalanceByID(w http.ResponseWriter, r *http.Request) {
	log := h.logger
	accountID := mux.Vars(r)["id"]

	balance, err := h.useCase.GetBalanceByID(r.Context(), accountID)
	if err != nil {
		var statusCode int
		var responseError server_http.Error
		switch {

		case errors.Is(err, accounts.ErrAccountNotFound):
			responseError = server_http.Error{Reason: accounts.ErrAccountNotFound.Error()}
			statusCode = http.StatusNotFound

		default:
			responseError = server_http.Error{Reason: "internal error server"}
			statusCode = http.StatusInternalServerError
		}
		log.WithFields(logrus.Fields{
			"status_code": statusCode,
		}).WithError(err).Error("get balance by id request failed")
		_ = server_http.SendResponse(w, responseError, statusCode)
		return
	}

	balanceResponse := GetBalanceByIdResponse{balance}
	_ = server_http.SendResponse(w, balanceResponse, http.StatusOK)
	log.WithFields(logrus.Fields{
		"account_id":  accountID,
		"status_code": http.StatusOK,
	}).Info("account balance found successfully")
}
