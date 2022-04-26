package accounts

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
	"github.com/sirupsen/logrus"
)

type Account struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	Balance   int    `json:"balance"`
}

type GetAllResponse struct {
	List []Account `json:"list"`
}

func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {

	log := h.logger
	accountsList, err := h.useCase.GetAll(r.Context())
	w.Header().Add("Content-Type", server_http.JSONContentType)
	if len(accountsList) == 0 && err != nil {
		log.WithError(err).Error("listing all accounts request failed")
		switch {
		case errors.Is(err, accounts.ErrAccountNotFound):
			w.WriteHeader(http.StatusNotFound)
			response := GetAllResponse{[]Account{}}
			json.NewEncoder(w).Encode(response)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			response := GetAllResponse{[]Account{}}
			json.NewEncoder(w).Encode(response)
		}
		return
	}

	getResponse := GetAllResponse{}
	for _, value := range accountsList {
		account := Account{value.ID, value.Name, value.CreatedAt.Format(server_http.DateLayout), value.Balance}
		getResponse.List = append(getResponse.List, account)
	}

	w.Header().Add("Content-Type", server_http.JSONContentType)
	responseGet := GetAllResponse{getResponse.List}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(responseGet)
	log.WithFields(logrus.Fields{
		"accounts_count": len(accountsList),
	}).Info("accounts listed successfully")

}
