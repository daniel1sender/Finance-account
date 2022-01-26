package accounts

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
	"github.com/gorilla/mux"
)

type ByIdResponse struct {
	Balance int `json:"balance"`
}

func (h Handler) GetBalanceByID(w http.ResponseWriter, r *http.Request) {

	accountID := mux.Vars(r)["id"]
	balance, err := h.useCase.GetBalanceByID(accountID)

	w.Header().Add("Content-Type", server_http.JSONContentType)
	if err != nil {
		log.Printf("get by id request failed: %s", err)
		switch {
		case errors.Is(err, accounts.ErrAccountNotFound):
			response := server_http.Error{Reason: accounts.ErrAccountNotFound.Error()}
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
}
