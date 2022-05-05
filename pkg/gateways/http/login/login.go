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

type LoginRequest struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	log := h.logger.WithFields(logrus.Fields{
		"route":  r.URL.Path,
		"method": r.Method,
	})


	var request LoginRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responseError := server_http.Error{Reason: "invalid request body"}
		_ = server_http.SendResponse(w, responseError, http.StatusBadRequest)
		log.WithFields(logrus.Fields{
			"status_code": http.StatusBadRequest,
		}).WithError(err).Error("error while decoding the body")
		return
	}

	token, err := h.UseCase.Login(r.Context(), request.Cpf, request.Secret)
	if err != nil {
		var responseError server_http.Error
		var statusCode int
		switch {
		case errors.Is(err, accounts.ErrAccountNotFound), errors.Is(err, login.ErrInvalidSecret):
			responseError = server_http.Error{Reason: login.ErrInvalidCredentials.Error()}
			statusCode = http.StatusForbidden

		case errors.Is(err, domain.ErrEmptySecret):
			responseError = server_http.Error{Reason: domain.ErrEmptySecret.Error()}
			statusCode = http.StatusBadRequest

		case errors.Is(err, domain.ErrInvalidCPF):
			responseError = server_http.Error{Reason: domain.ErrInvalidCPF.Error()}
			statusCode = http.StatusBadRequest

		default:
			responseError = server_http.Error{Reason: "internal server error"}
			statusCode = http.StatusInternalServerError
		}
		_ = server_http.SendResponse(w, responseError, statusCode)
		log.WithFields(logrus.Fields{
			"status_code": statusCode,
		}).WithError(err).Error("login request failed")
		return
	}

	var response = LoginResponse{Token: token}
	_ = server_http.SendResponse(w, response, http.StatusCreated)
	log.WithFields(logrus.Fields{
		"status_code": http.StatusCreated,
	}).Info("token was created successfully")
}
