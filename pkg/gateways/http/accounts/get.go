package accounts

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type Response struct {
	List []entities.Account
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request) {

	accountsList := h.useCase.Get()

	w.Header().Add("Content-Type", contentType)

	response := Response{accountsList}

	responseBody, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print("error while enconding the response")
	}

	_, err = w.Write(responseBody)
	if err != nil {
		log.Printf("error while getting the list of accounts")
	}
}
