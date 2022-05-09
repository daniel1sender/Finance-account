package accounts

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/daniel1sender/Desafio-API/pkg/domain"
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

	var createRequest CreateAccountRequest
	err := json.NewDecoder(r.Body).Decode(&createRequest)
	if err != nil {
		responseError := server_http.Error{Reason: "invalid request body"}
		_ = server_http.SendResponse(w, responseError, http.StatusBadRequest)
		log.WithFields(logrus.Fields{
			"status_code": http.StatusBadRequest,
		}).WithError(err).Error("error while decoding body")
		return
	}

	account, err := h.useCase.Create(r.Context(), createRequest.Name, createRequest.CPF, createRequest.Secret, createRequest.Balance)
	if err != nil {
		var responseError server_http.Error
		var statusCode int
		switch {

		case errors.Is(err, accounts.ErrExistingCPF):
			responseError = server_http.Error{Reason: accounts.ErrExistingCPF.Error()}
			statusCode = http.StatusConflict

		case errors.Is(err, entities.ErrInvalidName):
			responseError = server_http.Error{Reason: entities.ErrInvalidName.Error()}
			statusCode = http.StatusBadRequest

		case errors.Is(err, domain.ErrInvalidCPF):
			responseError = server_http.Error{Reason: domain.ErrInvalidCPF.Error()}
			statusCode = http.StatusBadRequest

		case errors.Is(err, domain.ErrEmptySecret):
			responseError = server_http.Error{Reason: domain.ErrEmptySecret.Error()}
			statusCode = http.StatusBadRequest

		case errors.Is(err, entities.ErrToGenerateHash):
			responseError = server_http.Error{Reason: entities.ErrToGenerateHash.Error()}
			w.WriteHeader(http.StatusInternalServerError)

		case errors.Is(err, entities.ErrNegativeBalance):
			responseError = server_http.Error{Reason: entities.ErrNegativeBalance.Error()}
			statusCode = http.StatusBadRequest

		default:
			responseError = server_http.Error{Reason: "internal server error"}
			statusCode = http.StatusInternalServerError
		}
		_ = server_http.SendResponse(w, responseError, statusCode)
		log.WithFields(logrus.Fields{
			"status_code": statusCode,
		}).WithError(err).Error("create account request failed")
		return
	}

	ExpectedCreateAt := account.CreatedAt.Format(server_http.DateLayout)
	CreateResponse := CreateAccountResponse{account.ID, account.Name, account.CPF, account.Balance, ExpectedCreateAt}
	_ = server_http.SendResponse(w, CreateResponse, http.StatusCreated)
	log.WithFields(logrus.Fields{
		"account_id": CreateResponse.ID,
		"status_code": http.StatusCreated,
	}).Info("successfully account created")

}
