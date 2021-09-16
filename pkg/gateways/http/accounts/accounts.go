package accounts

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
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
		log.Printf("failed to create an account: %s\n", err.Error())
		switch {

		case errors.Is(err, accounts.ErrExistingCPF):
			response := Error{Reason: accounts.ErrExistingCPF.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)

		case errors.Is(err, accounts.ErrCreatingNewAccount):
			response := Error{Reason: accounts.ErrCreatingNewAccount.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
		}
		return
	}

	createResponse := CreateResponse{account.ID, account.Name, account.CPF, account.Balance, account.CreatedAt}

	responseBody, _ := json.Marshal(createResponse)
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(responseBody)
	if err != nil {
		log.Printf("error while informing the new account")
	}
}
