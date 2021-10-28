package accounts

import (
	"encoding/json"
	"errors"
	"net/http"

	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/accounts"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ByIdResponse struct {
	Balance int `json:"balance"`
}

func (h Handler) GetBalanceByID(w http.ResponseWriter, r *http.Request) {
	log := h.logger
	requestID := r.Header.Get("Request-Id")
	if requestID == "" {
		log.Error("no request id informed")
		w.Header().Add("Content-Type", server_http.JSONContentType)
		response := server_http.Error{Reason: "invalid request header"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	accountID := mux.Vars(r)["id"]
	log = log.WithFields(logrus.Fields{
		"request_id": requestID,
		"account_id": accountID,
	})
	w.Header().Add("Content-Type", server_http.JSONContentType)

	balance, err := h.useCase.GetBalanceByID(accountID)
	if err != nil {
		log.Errorf("get balance by id request failed: %s", err)
		switch {
		case errors.Is(err, accounts.ErrIDNotFound):
			response := server_http.Error{Reason: accounts.ErrIDNotFound.Error()}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response)

		default:
			response := server_http.Error{Reason: "internal error server"}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)

		}
		return
	}

	balanceResponse := ByIdResponse{balance}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(balanceResponse)
	log.Info("account found successfully")
}
