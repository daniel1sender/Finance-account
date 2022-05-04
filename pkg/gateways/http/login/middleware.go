package login

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/daniel1sender/Desafio-API/pkg/domain/login"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
	"github.com/sirupsen/logrus"
)

var (
	ErrEmptyAuthHeader     = errors.New("got an empty authorization header")
	ErrInvalidHeaderFormat = errors.New("wrong authorization header format")
	ErrInvalidMethod       = errors.New("invalid authentication method")
)

func (h Handler) ValidateToken(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := h.logger
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response := server_http.Error{Reason: ErrEmptyAuthHeader.Error()}
			_ = server_http.SendResponse(w, response, http.StatusUnauthorized)
			log.WithFields(logrus.Fields{
				"status_code": http.StatusUnauthorized,
			}).WithError(ErrEmptyAuthHeader).Error("error occurred while was validating token")
			return
		}

		authString := strings.Split(authHeader, " ")
		if len(authString) != 2 {
			response := server_http.Error{Reason: ErrInvalidHeaderFormat.Error()}
			_ = server_http.SendResponse(w, response, http.StatusUnauthorized)
			log.WithFields(logrus.Fields{
				"status_code": http.StatusUnauthorized,
			}).WithError(ErrInvalidHeaderFormat).Error("error occurred while was validating token")
			return
		}

		if authString[0] != "Bearer" {
			response := server_http.Error{Reason: ErrInvalidMethod.Error()}
			_ = server_http.SendResponse(w, response, http.StatusUnauthorized)
			log.WithFields(logrus.Fields{
				"status_code": http.StatusUnauthorized,
			}).WithError(ErrInvalidMethod).Error("error occurred while was validating token")
			return
		}

		ctx := r.Context()
		claim, err := h.UseCase.ValidateToken(ctx, authString[1])
		if err != nil {
			switch {

			case errors.Is(err, login.ErrInvalidToken), errors.Is(err, login.ErrTokenNotFound):
				response := server_http.Error{Reason: login.ErrInvalidToken.Error()}
				_ = server_http.SendResponse(w, response, http.StatusForbidden)
				log.WithFields(logrus.Fields{
					"status_code": http.StatusForbidden,
				}).WithError(login.ErrInvalidToken).Error("error occurred while was validating token")

			default:
				response := server_http.Error{Reason: err.Error()}
				_ = server_http.SendResponse(w, response, http.StatusInternalServerError)
				log.WithFields(logrus.Fields{
					"status_code": http.StatusInternalServerError,
				}).WithError(err).Error("error occurred while was validating token")
			}
			return
		}

		accountID := claim.Sub
		ctx = context.WithValue(r.Context(), server_http.ContextAccountID, accountID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
