package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
)

type GetResponse struct {
	List []entities.Account `json:"list"`
}

func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {

	accountsList := h.useCase.GetAll()
	if len(accountsList) == 0 {
		w.Header().Add("Content-Type", server_http.JSONContentType)
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Add("Content-Type", server_http.JSONContentType)

	responseGet := GetResponse{accountsList}

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(responseGet)

}
