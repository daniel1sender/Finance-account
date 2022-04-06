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

func (h Handler) ValidateToken(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := h.logger
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			h.logger.Error("got an empty authorization header")
			response := server_http.Error{Reason: "got an empty authorization header"}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		authString := strings.Split(authHeader, " ")
		if len(authString) != 2 {
			h.logger.Error("wrong authorization header format")
			response := server_http.Error{Reason: "wrong authorization header format"}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		if authString[0] != "Bearer" {
			h.logger.Error("invalid authentication method")
			response := server_http.Error{Reason: "invalid authentication method"}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		ctx := r.Context()
		token, err := h.UseCase.ValidateToken(ctx, authString[1])
		if err != nil {
			log.WithError(err).Error("error occurred when was validating token")
			switch {
			case errors.Is(err, login.ErrInvalidToken):
				response := server_http.Error{Reason: login.ErrInvalidToken.Error()}
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(response)
			case errors.Is(err, login.ErrTokenNotFound):
				response := server_http.Error{Reason: login.ErrTokenNotFound.Error()}
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(response)
			default:
				response := server_http.Error{Reason: err.Error()}
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(response)
			}
			return
		}

		accountID := token.Sub
		ctx = context.WithValue(r.Context(), server_http.ContextAccountID, accountID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
