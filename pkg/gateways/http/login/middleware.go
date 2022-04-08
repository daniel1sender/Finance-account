package login

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/daniel1sender/Desafio-API/pkg/domain/login"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
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
			h.logger.Error(ErrEmptyAuthHeader)
			response := server_http.Error{Reason: ErrEmptyAuthHeader.Error()}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		authString := strings.Split(authHeader, " ")
		if len(authString) != 2 {
			h.logger.Error(ErrInvalidHeaderFormat)
			response := server_http.Error{Reason: ErrInvalidHeaderFormat.Error()}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		if authString[0] != "Bearer" {
			h.logger.Error(ErrInvalidMethod)
			response := server_http.Error{Reason: ErrInvalidMethod.Error()}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		ctx := r.Context()
		claim, err := h.UseCase.ValidateToken(ctx, authString[1])
		if err != nil {
			log.WithError(err).Error("error occurred while was validating token")
			switch {
			case errors.Is(err, login.ErrInvalidToken), errors.Is(err, login.ErrTokenNotFound):
				response := server_http.Error{Reason: login.ErrInvalidToken.Error()}
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(response)
			default:
				response := server_http.Error{Reason: err.Error()}
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(response)
			}
			return
		}

		accountID := claim.Sub
		ctx = context.WithValue(r.Context(), server_http.ContextAccountID, accountID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
