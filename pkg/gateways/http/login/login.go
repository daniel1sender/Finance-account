package login

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/daniel1sender/Desafio-API/pkg/domain"
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

	var statusCode int
	var request LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.Header().Add("Content-Type", server_http.JSONContentType)
		response := server_http.Error{Reason: "invalid request body"}
		_ = server_http.SendResponse(w, response, http.StatusBadRequest)
		log.WithFields(logrus.Fields{
			"status_code": http.StatusBadRequest,
		}).WithError(err).Error("error while decoding the body")
		return
	}

	w.Header().Add("Content-Type", server_http.JSONContentType)
	token, err := h.UseCase.Login(r.Context(), request.Cpf, request.Secret)
	if err != nil {
		switch {

		case errors.Is(err, accounts.ErrAccountNotFound), errors.Is(err, login.ErrInvalidSecret):
			response := server_http.Error{Reason: login.ErrInvalidCredentials.Error()}
			_ = server_http.SendResponse(w, response, http.StatusForbidden)
			statusCode = http.StatusForbidden

		case errors.Is(err, domain.ErrEmptySecret):
			response := server_http.Error{Reason: domain.ErrEmptySecret.Error()}
			_ = server_http.SendResponse(w, response, http.StatusBadRequest)
			statusCode = http.StatusBadRequest

		case errors.Is(err, domain.ErrInvalidCPF):
			response := server_http.Error{Reason: domain.ErrInvalidCPF.Error()}
			_ = server_http.SendResponse(w, response, http.StatusBadRequest)
			statusCode = http.StatusBadRequest

		default:
			response := server_http.Error{Reason: "internal server error"}
			_ = server_http.SendResponse(w, response, http.StatusInternalServerError)
			statusCode = http.StatusInternalServerError
		}
		log.WithFields(logrus.Fields{
			"status_code": statusCode,
		}).WithError(err).Error("login request failed")
		return
	}

	response := LoginUserResponse{Token: token}
	_ = server_http.SendResponse(w, response, http.StatusCreated)
	log.WithFields(logrus.Fields{
		"status_code": http.StatusCreated,
	}).Info("token was created successfully")
}
