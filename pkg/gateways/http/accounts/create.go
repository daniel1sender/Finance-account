package accounts

import (
	"encoding/json"
	"errors"
	"net/http"

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
	ID        string `json:"id"`
	Name      string `json:"name"`
	CPF       string `json:"cpf"`
	Balance   int    `json:"balance"`
	CreatedAt string `json:"created_at"`
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	log := h.logger
	requestID := r.Header.Get("Request-Id")
	if requestID == "" {
		log.Error("no request id informed")
		w.Header().Add("Content-Type", server_http.JSONContentType)
		response := server_http.Error{Reason: "invalid request header"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	log = log.WithField("request_id", requestID)

	var createRequest CreateRequest
	err := json.NewDecoder(r.Body).Decode(&createRequest)
	if err != nil {
		w.Header().Add("Content-Type", server_http.JSONContentType)
		response := server_http.Error{Reason: "invalid request body"}
		log.WithError(err).Errorf("error decoding body: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	account, err := h.useCase.Create(createRequest.Name, createRequest.CPF, createRequest.Secret, createRequest.Balance)
	w.Header().Add("Content-Type", server_http.JSONContentType)
	if err != nil {
		log.WithError(err).Error("create account request failed")
		switch {

		case errors.Is(err, accounts.ErrExistingCPF):
			response := server_http.Error{Reason: accounts.ErrExistingCPF.Error()}
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(response)

		case errors.Is(err, entities.ErrInvalidName):
			response := server_http.Error{Reason: entities.ErrInvalidName.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)

		case errors.Is(err, entities.ErrInvalidCPF):
			response := server_http.Error{Reason: entities.ErrInvalidCPF.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)

		case errors.Is(err, entities.ErrEmptySecret):
			response := server_http.Error{Reason: entities.ErrEmptySecret.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)

		case errors.Is(err, entities.ErrToGenerateHash):
			response := server_http.Error{Reason: entities.ErrToGenerateHash.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)

		case errors.Is(err, entities.ErrNegativeBalance):
			response := server_http.Error{Reason: entities.ErrNegativeBalance.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)

		default:
			response := server_http.Error{Reason: "internal server error"}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
		}
		return
	}

	ExpectedCreateAt := account.CreatedAt.Format(server_http.DateLayout)
	CreateResponse := CreateResponse{account.ID, account.Name, account.CPF, account.Balance, ExpectedCreateAt}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(CreateResponse)

}
