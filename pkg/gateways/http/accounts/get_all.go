package accounts

import (
	"encoding/json"
	"net/http"

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
	log := h.logger
	accountsList := h.useCase.GetAll()
	if len(accountsList) == 0 {
		w.Header().Add("Content-Type", server_http.JSONContentType)
		w.WriteHeader(http.StatusOK)
		log.Error("empty account list")
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
