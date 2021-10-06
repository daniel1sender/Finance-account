package accounts

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
)

type CreateRequest struct {
	Name    string `json:"name"`
	CPF     string `json:"cpf"`
	Secret  string `json:"secret"`
	Balance int    `json:"balance"`
}

type CreateResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CPF       string    `json:"cpf"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {

	var createRequest CreateRequest
	err := json.NewDecoder(r.Body).Decode(&createRequest)
	if err != nil {
		w.Header().Add("Content-Type", server_http.JSONContentType)
		response := Error{Reason: "invalid request body"}
		log.Printf("error decoding body: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	account, err := h.useCase.Create(createRequest.Name, createRequest.CPF, createRequest.Secret, createRequest.Balance)
	w.Header().Add("Content-Type", server_http.JSONContentType)
	if err != nil {
		log.Printf("request failed: %s\n", err.Error())
		switch {

		case errors.Is(err, accounts.ErrExistingCPF):
			response := Error{Reason: accounts.ErrExistingCPF.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)

		case errors.Is(err, entities.ErrInvalidName):
			response := Error{Reason: entities.ErrInvalidName.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)

		case errors.Is(err, entities.ErrInvalidCPF):
			response := Error{Reason: entities.ErrInvalidCPF.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)

		case errors.Is(err, entities.ErrEmptySecret):
			response := Error{Reason: entities.ErrEmptySecret.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)

		case errors.Is(err, entities.ErrToGenerateHash):
			response := Error{Reason: entities.ErrToGenerateHash.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)

		case errors.Is(err, entities.ErrNegativeBalance):
			response := Error{Reason: entities.ErrNegativeBalance.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)

		default:
			response := Error{Reason: "internal server error"}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
		}

		return
	}

	CreateResponse := CreateResponse{account.ID, account.Name, account.CPF, account.Balance, account.CreatedAt}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(CreateResponse)

}
