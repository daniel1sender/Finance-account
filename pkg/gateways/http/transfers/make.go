package transfers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
)

type Request struct {
	AccountOriginID      int `json:"account_origin_id"`
	AccountDestinationID int `json:"account_destination_id"`
	Amount               int `json:"amount"`
}

type Response struct {
	ID                   string    `json:"id"`
	AccountOriginID      int       `json:"account_origin_id"`
	AccountDestinationID int       `json:"account_destination_id"`
	Amount               int       `json:"amount"`
	CreatedAt            time.Time `json:"create_at"`
}

func (h Handler) Make(w http.ResponseWriter, r *http.Request) {

	var createRequest Request

	err := json.NewDecoder(r.Body).Decode(&createRequest)
	if err != nil {
		w.Header().Add("Content-Type", server_http.JSONContentType)
		log.Println("error != nil")
	}

	transfer, _ := h.useCase.Make(createRequest.AccountOriginID, createRequest.AccountDestinationID, createRequest.Amount)

	response := Response{transfer.ID, transfer.AccountOriginID, transfer.AccountDestinationID, transfer.Amount, transfer.CreatedAt}

	_ = json.NewEncoder(w).Encode(response)

}
