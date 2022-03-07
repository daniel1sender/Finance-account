package login

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
	"github.com/sirupsen/logrus"
)

type LoginRequest struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

type LoginResponse struct {
	AccountID string `json:"accountID"`
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	log := h.logger
	var requestBody LoginRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.Header().Add("Content-Type", server_http.JSONContentType)
		response := server_http.Error{Reason: "invalid request body"}
		log.WithError(err).Error("error decoding body")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Add("Content-Type", server_http.JSONContentType)
	token, accountID, err := h.useCase.Auth(r.Context(), requestBody.CPF, requestBody.Secret)
	if err != nil {
		log.WithError(err).Error("login request failed")
		switch {
		case errors.Is(err, accounts.ErrAccountNotFound):
			response := server_http.Error{Reason: accounts.ErrAccountNotFound.Error()}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
		default:
			response := server_http.Error{Reason: "internal error server"}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
		}
		return
	}
	w.Header().Add("Authorization", token)
	response := LoginResponse{accountID}
	_ = json.NewEncoder(w).Encode(response)
	log.WithFields(logrus.Fields{
		"sub": accountID,
	}).Info("account authenticated successfully")
}
