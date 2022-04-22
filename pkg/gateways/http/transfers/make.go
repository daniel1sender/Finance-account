package transfers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/domain/transfers/usecases"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
	"github.com/sirupsen/logrus"
)

type Request struct {
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

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	originAccountID := ctx.Value(server_http.ContextAccountID).(string)
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

	transfer, err := h.useCase.Create(ctx, originAccountID, createRequest.AccountDestinationID, createRequest.Amount)
	w.Header().Add("Content-Type", server_http.JSONContentType)
	if err != nil {
		log.WithError(err).Error("create transfer request failed")
		switch {

		case errors.Is(err, usecases.ErrOriginAccountNotFound):
			response := server_http.Error{Reason: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&response)

		case errors.Is(err, usecases.ErrDestinationAccountNotFound):
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

		case errors.Is(err, usecases.ErrInsufficientFunds):
			response := server_http.Error{Reason: usecases.ErrInsufficientFunds.Error()}
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
	response := Response{transfer.ID, originAccountID, transfer.AccountDestinationID, transfer.Amount, ExpectedCreateAt}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(response)
	log.WithFields(logrus.Fields{
		"origin_account_id":      originAccountID,
		"destination_account_id": transfer.AccountDestinationID,
	}).Info("transfer successful")
}
