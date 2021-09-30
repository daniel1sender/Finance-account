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

func (h Handler) Get(w http.ResponseWriter, r *http.Request) {

	accountsList := h.useCase.Get()
	if len(accountsList) == 0 {
		w.Header().Add("Content-Type", server_http.ContentType)
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Add("Content-Type", server_http.ContentType)

	responseGet := GetResponse{accountsList}

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(responseGet)

}
