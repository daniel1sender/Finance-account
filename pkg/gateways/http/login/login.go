package login

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/login"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
	"github.com/sirupsen/logrus"
)

type Request struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}

type Response struct {
	Token string `json:"token"`
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	log := h.logger
	log.WithFields(logrus.Fields{
		"route":  r.URL.Path,
		"method": r.Method,
	}).Info("login attempt realized")
	var request Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.Header().Add("Content-Type", server_http.JSONContentType)
		response := server_http.Error{Reason: "invalid request body"}
		log.WithError(err).Error("error decoding body")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	w.Header().Add("Content-Type", server_http.JSONContentType)
	token, err := h.UseCase.Login(r.Context(), request.Cpf, request.Secret)
	if err != nil {
		log.WithError(err).Error("login request failed")
		switch {
		case errors.Is(err, accounts.ErrAccountNotFound):
			response := server_http.Error{Reason: login.ErrInvalidCredetials.Error()}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
		case errors.Is(err, login.ErrEmptySecret):
			response := server_http.Error{Reason: login.ErrEmptySecret.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
		case errors.Is(err, login.ErrInvalidCPF):
			response := server_http.Error{Reason: login.ErrInvalidCPF.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
		case errors.Is(err, login.ErrInvalidSecret):
			response := server_http.Error{Reason: login.ErrInvalidSecret.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
		default:
			response := server_http.Error{Reason: "internal server error"}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
		}
		return
	}
	response := Response{Token: token}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(response)
	log.WithFields(logrus.Fields{
		"token": response.Token,
	}).Info("token created successfully")
}
