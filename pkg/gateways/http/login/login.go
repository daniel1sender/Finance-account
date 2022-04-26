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

type LoginUserRequest struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	log := h.logger.WithFields(logrus.Fields{
		"route":  r.URL.Path,
		"method": r.Method,
	})

	var request LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response := server_http.Error{Reason: "invalid request body"}
		log.WithError(err).Error("error while decoding the body")
		_ = server_http.Send(w, response, http.StatusBadRequest)
		return
	}

	token, err := h.UseCase.Login(r.Context(), request.Cpf, request.Secret)
	if err != nil {
		log.WithError(err).Error("login request failed")
		switch {

		case errors.Is(err, accounts.ErrAccountNotFound), errors.Is(err, login.ErrInvalidSecret):
			response := server_http.Error{Reason: login.ErrInvalidCredentials.Error()}
			_ = server_http.Send(w, response, http.StatusForbidden)

		case errors.Is(err, login.ErrEmptySecret):
			response := server_http.Error{Reason: login.ErrEmptySecret.Error()}
			_ = server_http.Send(w, response, http.StatusBadRequest)

		case errors.Is(err, login.ErrInvalidCPF):
			response := server_http.Error{Reason: login.ErrInvalidCPF.Error()}
			_ = server_http.Send(w, response, http.StatusBadRequest)
			
		default:
			response := server_http.Error{Reason: "internal server error"}
			_ = server_http.Send(w, response, http.StatusInternalServerError)
		}
		return
	}

	response := LoginUserResponse{Token: token}
	_ = server_http.Send(w, response, http.StatusCreated)
	log.Info("token was created successfully")
}
