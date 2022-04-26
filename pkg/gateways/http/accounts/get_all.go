package accounts

import (
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
	if len(accountsList) == 0 && err != nil {
		log.WithError(err).Error("listing all accounts request failed")
		switch {

		case errors.Is(err, accounts.ErrAccountNotFound):
			response := GetAllResponse{[]Account{}}
			_ = server_http.Send(w, response, http.StatusNotFound)

		default:
			response := GetAllResponse{[]Account{}}
			_ = server_http.Send(w, response, http.StatusInternalServerError)
		}
		return
	}

	response := GetAllResponse{}
	for _, value := range accountsList {
		account := Account{value.ID, value.Name, value.CreatedAt.Format(server_http.DateLayout), value.Balance}
		response.List = append(response.List, account)
	}

	responseList := GetAllResponse{response.List}
	_ = server_http.Send(w, responseList, http.StatusOK)
	log.WithFields(logrus.Fields{
		"accounts_count": len(accountsList),
	}).Info("accounts listed successfully")

}
