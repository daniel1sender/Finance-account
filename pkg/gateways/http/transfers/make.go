package transfers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/daniel1sender/Desafio-API/pkg/config"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/domain/login"
	"github.com/daniel1sender/Desafio-API/pkg/domain/transfers"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
	"github.com/sirupsen/logrus"
)

type Request struct {
	AccountOriginID      string `json:"account_origin_id"`
	Token                string `json:"token"`
	AccountDestinationID string `json:"account_destination_id"`
	Amount               int    `json:"amount"`
}

type Response struct {
	ID                   string `json:"id"`
	AccountOriginID      string `json:"account_origin_id"`
	AccountDestinationID string `json:"account_destination_id"`
	Amount               int    `json:"amount"`
	CreatedAt            string `json:"create_at"`
}

func (h Handler) Make(w http.ResponseWriter, r *http.Request) {
	log := h.logger
	var createRequest Request
	err := json.NewDecoder(r.Body).Decode(&createRequest)
	if err != nil {
		w.Header().Add("Content-Type", server_http.JSONContentType)
		response := server_http.Error{Reason: "invalid request body"}
		log.WithError(err).Error("error decoding body")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	config, err := config.GetConfig()
	if err != nil {
		w.Header().Add("Content-Type", server_http.JSONContentType)
		response := server_http.Error{Reason: "unable to get environment variables"}
		log.WithError(err).Error("unable to get environment variables")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = h.loginUseCase.CheckToken(r.Context(), createRequest.Token)
	if err != nil {
		w.Header().Add("Content-Type", server_http.JSONContentType)
		response := server_http.Error{Reason: "token not found"}
		log.WithError(err).Error("error while searching for token")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = login.ValidateToken(createRequest.Token, createRequest.AccountOriginID, config.TokenSecret)
	if err != nil {
		w.Header().Add("Content-Type", server_http.JSONContentType)
		response := server_http.Error{Reason: "unable to validate token"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		log.WithError(err).Error("error while validating token")
		return
	}

	transfer, err := h.useCase.Make(r.Context(), createRequest.AccountOriginID, createRequest.AccountDestinationID, createRequest.Amount)
	w.Header().Add("Content-Type", server_http.JSONContentType)
	if err != nil {
		log.WithError(err).Error("create transfer request failed")
		switch {

		case errors.Is(err, transfers.ErrOriginAccountNotFound):
			response := server_http.Error{Reason: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&response)

		case errors.Is(err, transfers.ErrDestinationAccountNotFound):
			response := server_http.Error{Reason: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&response)

		case errors.Is(err, entities.ErrAmountLessOrEqualZero):
			response := server_http.Error{Reason: entities.ErrAmountLessOrEqualZero.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&response)

		case errors.Is(err, entities.ErrSameAccountTransfer):
			response := server_http.Error{Reason: entities.ErrSameAccountTransfer.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&response)

		case errors.Is(err, transfers.ErrInsufficientFunds):
			response := server_http.Error{Reason: transfers.ErrInsufficientFunds.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&response)

		default:
			response := server_http.Error{Reason: "internal server error"}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&response)
		}
		return
	}

	ExpectedCreateAt := transfer.CreatedAt.Format(server_http.DateLayout)
	response := Response{transfer.ID, transfer.AccountOriginID, transfer.AccountDestinationID, transfer.Amount, ExpectedCreateAt}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(response)
	log.WithFields(logrus.Fields{
		"origin_account_id":      transfer.AccountOriginID,
		"destination_account_id": transfer.AccountDestinationID,
	}).Info("transfer successful")
}
