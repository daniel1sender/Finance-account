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

type CreateTransferRequest struct {
	AccountDestinationID string `json:"account_destination_id"`
	Amount               int    `json:"amount"`
}

type CreateTransferResponse struct {
	ID                   string `json:"id"`
	AccountOriginID      string `json:"account_origin_id"`
	AccountDestinationID string `json:"account_destination_id"`
	Amount               int    `json:"amount"`
	CreatedAt            string `json:"create_at"`
}

func (h Handler) Make(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	originAccountID := ctx.Value(server_http.ContextAccountID).(string)
	log := h.logger
	var request CreateTransferRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response := server_http.Error{Reason: "invalid request body"}
		log.WithError(err).Error("error decoding body")
		_ = server_http.Send(w, response, http.StatusBadRequest)
		return
	}

	transfer, err := h.useCase.Make(ctx, originAccountID, request.AccountDestinationID, request.Amount)
	if err != nil {
		log.WithError(err).Error("create transfer request failed")
		switch {

		case errors.Is(err, usecases.ErrOriginAccountNotFound):
			response := server_http.Error{Reason: err.Error()}
			_ = server_http.Send(w, response, http.StatusBadRequest)

		case errors.Is(err, usecases.ErrDestinationAccountNotFound):
			response := server_http.Error{Reason: err.Error()}
			_ = server_http.Send(w, response, http.StatusBadRequest)

		case errors.Is(err, entities.ErrAmountLessOrEqualZero):
			response := server_http.Error{Reason: entities.ErrAmountLessOrEqualZero.Error()}
			_ = server_http.Send(w, response, http.StatusBadRequest)

		case errors.Is(err, entities.ErrSameAccountTransfer):
			response := server_http.Error{Reason: entities.ErrSameAccountTransfer.Error()}
			_ = server_http.Send(w, response, http.StatusBadRequest)

		case errors.Is(err, usecases.ErrInsufficientFunds):
			response := server_http.Error{Reason: usecases.ErrInsufficientFunds.Error()}
			_ = server_http.Send(w, response, http.StatusBadRequest)

		default:
			response := server_http.Error{Reason: "internal server error"}
			_ = server_http.Send(w, response, http.StatusInternalServerError)
		}
		return
	}

	ExpectedCreateAt := transfer.CreatedAt.Format(server_http.DateLayout)
	response := CreateTransferResponse{transfer.ID, originAccountID, transfer.AccountDestinationID, transfer.Amount, ExpectedCreateAt}
	_ = server_http.Send(w, response, http.StatusCreated)
	log.WithFields(logrus.Fields{
		"origin_account_id":      originAccountID,
		"destination_account_id": transfer.AccountDestinationID,
	}).Info("transfer successful")
}
