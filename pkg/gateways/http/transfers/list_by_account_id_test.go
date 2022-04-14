package transfers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/domain/transfers"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestHandlerListByAccountID(t *testing.T) {
	log := logrus.NewEntry(logrus.New())

	t.Run("should return 200 and the list of transfers", func(t *testing.T) {
		transfer := entities.Transfer{AccountOriginID: "1", AccountDestinationID: "0", Amount: 10}
		transferList := make([]entities.Transfer, 0)
		transferList = append(transferList, transfer)
		useCase := transfers.UseCaseMock{ListTransfers: transferList}
		h := NewHandler(&useCase, log)
		newRequest, _ := http.NewRequest(http.MethodGet, "/transfers", nil)
		ctx := context.WithValue(newRequest.Context(), server_http.ContextAccountID, transfer.AccountOriginID)
		newResponse := httptest.NewRecorder()
		h.ListByAccountID(newResponse, newRequest.WithContext(ctx))

		var response []ResponseList
		_ = json.Unmarshal(newResponse.Body.Bytes(), &response)

		assert.Equal(t, newResponse.Header().Get("content-type"), server_http.JSONContentType)
		assert.Equal(t, newResponse.Code, http.StatusOK)
		assert.NotEmpty(t, response)
	})

	t.Run("should return 404 and an error message when there are no transfers from that account", func(t *testing.T) {
		useCase := transfers.UseCaseMock{Error: transfers.ErrEmptyList}
		h := NewHandler(&useCase, log)
		newRequest, _ := http.NewRequest(http.MethodGet, "/transfers", nil)
		ctx := context.WithValue(newRequest.Context(), server_http.ContextAccountID, "123")
		newResponse := httptest.NewRecorder()
		h.ListByAccountID(newResponse, newRequest.WithContext(ctx))

		var response server_http.Error
		_ = json.Unmarshal(newResponse.Body.Bytes(), &response)

		assert.Equal(t, newResponse.Header().Get("content-type"), server_http.JSONContentType)
		assert.Equal(t, newResponse.Code, http.StatusNotFound)
		assert.Equal(t, response.Reason, transfers.ErrTransfersNotFound.Error())
	})

	t.Run("should return 500 and an error when an unexpected error occourred", func(t *testing.T) {
		unexpectedError := errors.New("unexpected error")
		useCase := transfers.UseCaseMock{Error: unexpectedError}
		h := NewHandler(&useCase, log)
		newRequest, _ := http.NewRequest(http.MethodGet, "/transfers", nil)
		ctx := context.WithValue(newRequest.Context(), server_http.ContextAccountID, "123")
		newResponse := httptest.NewRecorder()
		h.ListByAccountID(newResponse, newRequest.WithContext(ctx))

		var response server_http.Error
		_ = json.Unmarshal(newResponse.Body.Bytes(), &response)

		assert.Equal(t, newResponse.Header().Get("content-type"), server_http.JSONContentType)
		assert.Equal(t, newResponse.Code, http.StatusInternalServerError)
		assert.Equal(t, unexpectedError.Error(), response.Reason)
	})
}
