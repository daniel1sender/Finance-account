package transfers

import (
	"encoding/json"
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

func (h Handler) ListByID(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	originAccountID := ctx.Value(server_http.ContextAccountID).(string)
	log := h.logger

	transfersList, err := h.useCase.ListByAccountID(ctx, originAccountID)
	w.Header().Add("Content-Type", server_http.JSONContentType)
	if err != nil {
		log.WithError(err).Error("transfers listing request failed")
		switch {
		case errors.Is(err, transfers.ErrEmptyList):
			response := server_http.Error{Reason: transfers.ErrTransfersNotFound.Error()}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
		default:
			response := server_http.Error{Reason: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
		}
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

	w.WriteHeader(http.StatusOK)
	log.WithFields(logrus.Fields{
		"origin_account_id":      originAccountID,
		"total_transfers_listed": len(transfersList),
	}).Info("transfers listed successfully")
	_ = json.NewEncoder(w).Encode(response)
}
