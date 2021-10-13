package transfers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
)

type Request struct {
	AccountOriginID      string `json:"account_origin_id"`
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

	var createRequest Request
	err := json.NewDecoder(r.Body).Decode(&createRequest)
	if err != nil {
		w.Header().Add("Content-Type", server_http.JSONContentType)
		response := server_http.Error{Reason: "invalid request body"}
		log.Printf("error decoding body: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	transfer, err := h.useCase.Make(createRequest.AccountOriginID, createRequest.AccountDestinationID, createRequest.Amount)
	w.Header().Add("Content-Type", server_http.JSONContentType)
	if err != nil {
		log.Printf("create transfer request failed: %s\n", err.Error())
		switch {

		case errors.Is(err, entities.ErrAmountLessOrEqualZero):
			response := server_http.Error{Reason: entities.ErrAmountLessOrEqualZero.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&response)

		case errors.Is(err, entities.ErrSameAccountTransfer):
			response := server_http.Error{Reason: entities.ErrSameAccountTransfer.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&response)

		default:
			response := server_http.Error{Reason: "internal server error"}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&response)
		}

		return

	}

	createAt := transfer.CreatedAt
	ExpectedCreateAt := createAt.Format(server_http.DateLayout)
	response := Response{transfer.ID, transfer.AccountOriginID, transfer.AccountDestinationID, transfer.Amount, ExpectedCreateAt}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(response)

}
