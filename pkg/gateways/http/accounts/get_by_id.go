package accounts

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	server_http "github.com/daniel1sender/Desafio-API/pkg/domain/gateways/http"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/accounts"
)

type RequestByID struct {
	Id string `json:"id"`
}

type ResponseByID struct {
	Balance int `json:"balance"`
}

func (h Handler) GetBalanceByID(w http.ResponseWriter, r *http.Request) {

	var request RequestByID
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response := Error{Reason: "invalid request body"}
		log.Printf("error decoding body: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	balance, err := h.useCase.GetBalanceByID(request.Id)
	w.Header().Add("Content-Type", server_http.ContentType)
	if err != nil {
		log.Printf("request failed: %s", err)
		switch {
		case errors.Is(err, accounts.ErrIDNotFound):
			response := Error{Reason: accounts.ErrIDNotFound.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)

		default:
			response := Error{Reason: "internal error server"}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
		}
		return
	}

	balanceResponse := ResponseByID{balance}

	response, err := json.Marshal(balanceResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error while enconding the response")
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(response)
	if err != nil {
		log.Printf("error while informing the new account")
		return
	}

}
