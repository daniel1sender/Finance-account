package accounts

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

const (
	contentType = "application/json"
)

type Handler struct {
	useCase accounts.AccountUseCase
}

func NewHandler(useCase accounts.AccountUseCase) Handler {
	return Handler{
		useCase: useCase,
	}
}

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

type Error struct {
	Reason string `json:"reason"`
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {

	var createRequest CreateRequest
	err := json.NewDecoder(r.Body).Decode(&createRequest)
	if err != nil {
		response := Error{Reason: "invalid request body"}
		log.Printf("error decoding body: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	account, err := h.useCase.Create(createRequest.Name, createRequest.CPF, createRequest.Secret, createRequest.Balance)
	w.Header().Add("Content-Type", contentType)
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

		case errors.Is(err, entities.ErrBlankSecret):
			response := Error{Reason: entities.ErrBlankSecret.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)

		case errors.Is(err, entities.ErrToGenerateHash):
			response := Error{Reason: entities.ErrToGenerateHash.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)

		case errors.Is(err, entities.ErrBalanceLessZero):
			response := Error{Reason: entities.ErrBalanceLessZero.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)

		default:
			response := Error{Reason: "internal server error"}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
		}

		return
	}

	createResponse := CreateResponse{account.ID, account.Name, account.CPF, account.Balance, account.CreatedAt}

	responseBody, err := json.Marshal(createResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print("error while enconding the response")
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(responseBody)
	if err != nil {
		log.Printf("error while informing the new account")
	}
}
