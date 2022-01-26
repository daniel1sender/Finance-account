package accounts

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
)

type Account struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	Balance   int    `json:"balance"`
}

type GetResponse struct {
	List []Account `json:"list"`
}

func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {

	accountsList, err := h.useCase.GetAll()
	w.Header().Add("Content-Type", server_http.JSONContentType)
	if len(accountsList) == 0 && err != nil {
		log.Printf("get all request failed: %s", err)
		switch {
		case errors.Is(err, accounts.ErrAccountNotFound):
			w.WriteHeader(http.StatusNotFound)
			response := GetResponse{[]Account{}}
			json.NewEncoder(w).Encode(response)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			response := GetResponse{[]Account{}}
			json.NewEncoder(w).Encode(response)
		}
		return
	}

	getResponse := GetResponse{}
	for _, value := range accountsList {
		account := Account{value.ID, value.Name, value.CreatedAt.Format(server_http.DateLayout), value.Balance}
		getResponse.List = append(getResponse.List, account)
	}

	w.Header().Add("Content-Type", server_http.JSONContentType)
	responseGet := GetResponse{getResponse.List}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(responseGet)

}
