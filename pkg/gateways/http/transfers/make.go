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

type TransferRequest struct {
	AccountDestinationID string `json:"account_destination_id"`
	Amount               int    `json:"amount"`
}

type TransferResponse struct {
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
	var statusCode int
	var createRequest TransferRequest
	err := json.NewDecoder(r.Body).Decode(&createRequest)
	if err != nil {
		response := server_http.Error{Reason: "invalid request body"}
		statusCode = http.StatusBadRequest
		_ = server_http.SendResponse(w, response, statusCode)
		log.WithFields(logrus.Fields{
			"status_code": statusCode,
		}).WithError(err).Error("error while decoding body")
		return
	}

	transfer, err := h.useCase.Make(ctx, originAccountID, createRequest.AccountDestinationID, createRequest.Amount)
	if err != nil {
		var responseError server_http.Error
		switch {

		case errors.Is(err, usecases.ErrOriginAccountNotFound):
			responseError = server_http.Error{Reason: err.Error()}
			statusCode = http.StatusBadRequest

		case errors.Is(err, usecases.ErrDestinationAccountNotFound):
			responseError = server_http.Error{Reason: err.Error()}
			statusCode = http.StatusBadRequest

		case errors.Is(err, entities.ErrAmountLessOrEqualZero):
			responseError = server_http.Error{Reason: entities.ErrAmountLessOrEqualZero.Error()}
			statusCode = http.StatusBadRequest

		case errors.Is(err, entities.ErrSameAccountTransfer):
			responseError = server_http.Error{Reason: entities.ErrSameAccountTransfer.Error()}
			statusCode = http.StatusBadRequest

		case errors.Is(err, usecases.ErrInsufficientFunds):
			responseError = server_http.Error{Reason: usecases.ErrInsufficientFunds.Error()}
			statusCode = http.StatusBadRequest

		default:
			responseError = server_http.Error{Reason: "internal server error"}
			statusCode = http.StatusInternalServerError
		}
		_ = server_http.SendResponse(w, responseError, statusCode)
		log.WithFields(logrus.Fields{
			"status_code": statusCode,
		}).WithError(err).Error("make transfer request failed")
		return
	}

	ExpectedCreateAt := transfer.CreatedAt.Format(server_http.DateLayout)
	response := TransferResponse{transfer.ID, originAccountID, transfer.AccountDestinationID, transfer.Amount, ExpectedCreateAt}
	statusCode = http.StatusCreated
	_ = server_http.SendResponse(w, response, statusCode)
	log.WithFields(logrus.Fields{
		"origin_account_id":      originAccountID,
		"destination_account_id": transfer.AccountDestinationID,
		"status_code":            statusCode,
	}).Info("transfer successful")
}
