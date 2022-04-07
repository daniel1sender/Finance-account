package login

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/domain/login"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestMiddlewareValidateToken(t *testing.T) {
	log := logrus.NewEntry(logrus.New())

	t.Run("should validate the token", func(t *testing.T) {
		type responseBody struct {
			Reason string
		}
		header := http.Header{
			"Authorization": []string{"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxZDlhY2RhNS1lZTRlLTQ5ZDMtOTdiMy1hN2UxYmE5ZWNiNmMiLCJleHAiOjE2NDkxNzU5NzAsImlhdCI6MTY0OTE3NTY3MCwianRpIjoiNzJhNGExYWItNmI4Zi00NDRmLTg4ODYtZjc1NDI0NjYwZGFjIn0.DPE5rEofHuWpSly0qE_5kZ__EvPQP_vySNTxT3FJ7A0"},
		}
		useCase := login.UseCaseMock{
			Claims: entities.Claims{Sub: "95571638-1aa3-4ee8-a05f-4c2d6c97ef4e"},
			Token:  "",
			Error:  nil,
		}
		handler := NewHandler(&useCase, log)
		request, _ := http.NewRequest(http.MethodGet, "/login", nil)
		request.Header = header
		newResponse := httptest.NewRecorder()
		validate := handler.ValidateToken(http.HandlerFunc(createHandleFunc))
		validate(newResponse, request)
		var response responseBody
		json.Unmarshal(newResponse.Body.Bytes(), &response)

		got := newResponse.Body.String()
		expected := useCase.Claims.Sub

		assert.Equal(t, newResponse.Code, http.StatusOK)
		assert.Equal(t, expected, got)

	})

	t.Run("should return 400 when the authorization header is empty", func(t *testing.T) {
		header := http.Header{
			"Authorization": []string{},
		}
		useCase := login.UseCaseMock{}
		handler := NewHandler(&useCase, log)
		request, _ := http.NewRequest(http.MethodGet, "/login", nil)
		newResponse := httptest.NewRecorder()
		request.Header = header
		validate := handler.ValidateToken(http.HandlerFunc(createHandleFunc))
		validate(newResponse, request)
		var response server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &response)

		got := response.Reason
		expected := "got an empty authorization header"

		assert.Equal(t, http.StatusUnauthorized, newResponse.Code)
		assert.Equal(t, expected, got)
	})

	t.Run("should return 400 when got a wrong authorization header format", func(t *testing.T) {
		header := http.Header{
			"Authorization": []string{"Bearer"},
		}
		useCase := login.UseCaseMock{}
		handler := NewHandler(&useCase, log)
		request, _ := http.NewRequest(http.MethodGet, "/login", nil)
		newResponse := httptest.NewRecorder()
		request.Header = header
		validate := handler.ValidateToken(http.HandlerFunc(createHandleFunc))
		validate(newResponse, request)

		var response server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &response)

		got := response.Reason
		expected := "wrong authorization header format"

		assert.Equal(t, http.StatusUnauthorized, newResponse.Code)
		assert.Equal(t, expected, got)
	})

	t.Run("should return 400 when the authentication method is wrong", func(t *testing.T) {
		header := http.Header{
			"Authorization": []string{"method token"},
		}
		useCase := login.UseCaseMock{}
		handler := NewHandler(&useCase, log)
		request, _ := http.NewRequest(http.MethodGet, "/login", nil)
		newResponse := httptest.NewRecorder()
		request.Header = header
		validate := handler.ValidateToken(http.HandlerFunc(createHandleFunc))
		validate(newResponse, request)

		var response server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &response)

		got := response.Reason
		expected := "invalid authentication method"

		assert.Equal(t, http.StatusUnauthorized, newResponse.Code)
		assert.Equal(t, expected, got)
	})

	t.Run("should return 403 when the token is invalid", func(t *testing.T) {
		header := http.Header{
			"Authorization": []string{"Bearer "},
		}
		useCase := login.UseCaseMock{Error: login.ErrInvalidToken}
		handler := NewHandler(&useCase, log)
		request, _ := http.NewRequest(http.MethodGet, "/login", nil)
		newResponse := httptest.NewRecorder()
		request.Header = header
		validate := handler.ValidateToken(http.HandlerFunc(createHandleFunc))
		validate(newResponse, request)

		var response server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &response)

		got := response.Reason
		expected := login.ErrInvalidToken.Error()

		assert.Equal(t, http.StatusForbidden, newResponse.Code)
		assert.Equal(t, expected, got)
	})

	t.Run("should return 404 when the token is not found", func(t *testing.T) {
		header := http.Header{
			"Authorization": []string{"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxZDlhY2RhNS1lZTRlLTQ5ZDMtOTdiMy1hN2UxYmE5ZWNiNmMiLCJleHAiOjE2NDkxNzU5NzAsImlhdCI6MTY0OTE3NTY3MCwianRpIjoiNzJhNGExYWItNmI4Zi00NDRmLTg4ODYtZjc1NDI0NjYwZGFjIn0.DPE5rEofHuWpSly0qE_5kZ__EvPQP_vySNTxT3FJ7A0"},
		}
		useCase := login.UseCaseMock{Error: login.ErrTokenNotFound}
		handler := NewHandler(&useCase, log)
		request, _ := http.NewRequest(http.MethodGet, "/login", nil)
		newResponse := httptest.NewRecorder()
		request.Header = header
		validate := handler.ValidateToken(http.HandlerFunc(createHandleFunc))
		validate(newResponse, request)

		var response server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &response)

		got := response.Reason
		expected := login.ErrTokenNotFound.Error()

		assert.Equal(t, http.StatusForbidden, newResponse.Code)
		assert.Equal(t, expected, got)
	})

	t.Run("should return 500 and an error when an unexpected error occurred", func(t *testing.T) {
		header := http.Header{
			"Authorization": []string{"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxZDlhY2RhNS1lZTRlLTQ5ZDMtOTdiMy1hN2UxYmE5ZWNiNmMiLCJleHAiOjE2NDkxNzU5NzAsImlhdCI6MTY0OTE3NTY3MCwianRpIjoiNzJhNGExYWItNmI4Zi00NDRmLTg4ODYtZjc1NDI0NjYwZGFjIn0.DPE5rEofHuWpSly0qE_5kZ__EvPQP_vySNTxT3FJ7A0"},
		}
		unexpectedError := errors.New("unexpected error")
		useCase := login.UseCaseMock{Error: unexpectedError}
		handler := NewHandler(&useCase, log)
		request, _ := http.NewRequest(http.MethodGet, "/login", nil)
		newResponse := httptest.NewRecorder()
		request.Header = header
		validate := handler.ValidateToken(http.HandlerFunc(createHandleFunc))
		validate(newResponse, request)

		var response server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &response)

		got := response.Reason
		expected := unexpectedError.Error()

		assert.Equal(t, http.StatusInternalServerError, newResponse.Code)
		assert.Equal(t, expected, got)
	})
}

func createHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	accountID := r.Context().Value(server_http.ContextAccountID)
	w.Write([]byte(accountID.(string)))
}
