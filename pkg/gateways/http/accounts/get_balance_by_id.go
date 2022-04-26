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
		log.WithError(err).Error("get balance by id request failed")
		switch {
			
		case errors.Is(err, accounts.ErrAccountNotFound):
			response := server_http.Error{Reason: accounts.ErrAccountNotFound.Error()}
			_ = server_http.Send(w, response, http.StatusNotFound)

		default:
			response := server_http.Error{Reason: "internal error server"}
			_ = server_http.Send(w, response, http.StatusInternalServerError)
		}
		return
	}

	response := GetBalanceByIdResponse{balance}
	_ = server_http.Send(w, response, http.StatusOK)
	log.WithFields(logrus.Fields{
		"account_id": accountID,
	}).Info("account balance found successfully")
}
