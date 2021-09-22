package accounts

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
)

type ResponseGet struct {
	List []entities.Account
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request) {

	accountsList := h.useCase.Get()
	if len(accountsList) == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Add("Content-Type", server_http.ContentType)

	responseGet := ResponseGet{accountsList}

	response, err := json.Marshal(responseGet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("JSON marshaling failed: %s", err)
	}

	_, err = w.Write(response)
	if err != nil {
		log.Printf("error while getting the list of accounts")
	}
}
