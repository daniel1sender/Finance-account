package transfers

import (
	"errors"
	"net/http"

	"github.com/daniel1sender/Desafio-API/pkg/domain/transfers"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
	"github.com/sirupsen/logrus"
)

type ResponseList struct {
	TransferID           string `json:"id"`
	AccountOriginID      string `json:"account_origin_id"`
	AccountDestinationID string `json:"account_destination_id"`
	Amount               int    `json:"amount"`
	CreatedAt            string `json:"created_at"`
}

func (h Handler) ListByAccountID(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	var statusCode int
	originAccountID := ctx.Value(server_http.ContextAccountID).(string)
	log := h.logger.WithFields(logrus.Fields{
		"origin_account_id": originAccountID,
	})

	transfersList, err := h.useCase.ListByAccountID(ctx, originAccountID)
	if err != nil {
		var responseError server_http.Error
		switch {

		case errors.Is(err, transfers.ErrEmptyList):
			responseError = server_http.Error{Reason: transfers.ErrEmptyList.Error()}
			statusCode = http.StatusNotFound
			
		default:
			responseError = server_http.Error{Reason: err.Error()}
			statusCode = http.StatusInternalServerError
		}
		_ = server_http.SendResponse(w, responseError, statusCode)
		log.WithFields(logrus.Fields{
			"status_code": statusCode,
		}).WithError(err).Error("transfers listing request failed")
		return
	}

	response := make([]ResponseList, len(transfersList))
	for index, transfer := range transfersList {
		response[index].TransferID = transfer.ID
		response[index].AccountOriginID = transfer.AccountOriginID
		response[index].AccountDestinationID = transfer.AccountDestinationID
		response[index].Amount = transfer.Amount
		response[index].CreatedAt = transfer.CreatedAt.String()
	}

	statusCode = http.StatusOK
	_ = server_http.SendResponse(w, response, statusCode)
	log.WithFields(logrus.Fields{
		"total_transfers_listed": len(transfersList),
		"status_code":            statusCode,
	}).Info("transfers listed successfully")
}
