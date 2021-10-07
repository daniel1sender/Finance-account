package transfers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	transfers_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/transfers"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
	transfers_storage "github.com/daniel1sender/Desafio-API/pkg/gateways/store/transfers"
)

func TestMake(t *testing.T) {

	t.Run("should return 201 and null error when the type informed is a json", func(t *testing.T) {

		createRequest := Request{AccountOriginID: 1, AccountDestinationID: 0, Amount: 10}

		storage := transfers_storage.NewStorage()
		useCase := transfers_usecase.NewTransferUseCase(storage)
		h := NewHandler(useCase)

		request, _ := json.Marshal(createRequest)

		newRequest, _ := http.NewRequest(http.MethodPost, "/transfers", bytes.NewReader(request))
		newResponse := httptest.NewRecorder()

		h.Make(newResponse, newRequest)

		var response Response

		_ = json.Unmarshal(newResponse.Body.Bytes(), &response)

		if newResponse.Code != http.StatusCreated {
			t.Errorf("expected '%d' but got '%d'", http.StatusCreated, newResponse.Code)
		}

		headerType := newResponse.Header().Get("content-type")
		if headerType != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, headerType)
		}

		if response.AccountOriginID != createRequest.AccountOriginID {
			t.Errorf("expected '%d' but got '%d'", createRequest.AccountOriginID, response.AccountOriginID)
		}

		if response.AccountDestinationID != createRequest.AccountDestinationID {
			t.Errorf("expected '%d' but got '%d'", createRequest.AccountDestinationID, response.AccountDestinationID)
		}

		if response.Amount != createRequest.Amount {
			t.Errorf("expected '%d' but got '%d'", createRequest.Amount, response.Amount)
		}

		if response.CreatedAt.IsZero() {
			t.Errorf("expected nonzero time but got '%s'", response.CreatedAt)
		}

	})

	/* 	t.Run("should return 400 when amount is less or equal zero", func(t 		createRequest := Request{AccountOriginID: 1, *testing.T) {

	AccountDestinationID: 0, Amount: 10}

			storage := transfers_storage.NewStorage()
			useCase := transfers_usecase.NewTransferUseCase(storage)
			h := NewHandler(useCase)

		}) */

}
