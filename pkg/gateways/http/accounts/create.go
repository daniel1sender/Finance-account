package accounts

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
	"github.com/sirupsen/logrus"
)

type CreateAccountRequest struct {
	Name    string `json:"name"`
	CPF     string `json:"cpf"`
	Secret  string `json:"secret"`
	Balance int    `json:"balance"`
}

type CreateAccountResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CPF       string `json:"cpf"`
	Balance   int    `json:"balance"`
	CreatedAt string `json:"created_at"`
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {

	log := h.logger

	var request CreateAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.WithError(err).Error("error decoding body")
		response := server_http.Error{Reason: "invalid request body"}
		_ = server_http.Send(w, response, http.StatusBadRequest)
		return
	}

	account, err := h.useCase.Create(r.Context(), request.Name, request.CPF, request.Secret, request.Balance)
	if err != nil {
		log.WithError(err).Error("create account request failed")
		switch {
		case errors.Is(err, accounts.ErrExistingCPF):
			response := server_http.Error{Reason: accounts.ErrExistingCPF.Error()}
			_ = server_http.Send(w, response, http.StatusConflict)

		case errors.Is(err, entities.ErrInvalidName):
			response := server_http.Error{Reason: entities.ErrInvalidName.Error()}
			_ = server_http.Send(w, response, http.StatusBadRequest)

		case errors.Is(err, entities.ErrInvalidCPF):
			response := server_http.Error{Reason: entities.ErrInvalidCPF.Error()}
			_ = server_http.Send(w, response, http.StatusBadRequest)

		case errors.Is(err, entities.ErrEmptySecret):
			response := server_http.Error{Reason: entities.ErrEmptySecret.Error()}
			_ = server_http.Send(w, response, http.StatusBadRequest)

		case errors.Is(err, entities.ErrToGenerateHash):
			response := server_http.Error{Reason: entities.ErrToGenerateHash.Error()}
			_ = server_http.Send(w, response, http.StatusInternalServerError)

		case errors.Is(err, entities.ErrNegativeBalance):
			response := server_http.Error{Reason: entities.ErrNegativeBalance.Error()}
			_ = server_http.Send(w, response, http.StatusBadRequest)

		default:
			response := server_http.Error{Reason: err.Error()}
			_ = server_http.Send(w, response, http.StatusInternalServerError)
		}
		return
	}

	expectedCreateAt := account.CreatedAt.Format(server_http.DateLayout)
	response := CreateAccountResponse{account.ID, account.Name, account.CPF, account.Balance, expectedCreateAt}
	_ = server_http.Send(w, response, http.StatusCreated)
	log.WithFields(logrus.Fields{
		"account_id": response.ID,
	}).Info("successfully account created")
}
