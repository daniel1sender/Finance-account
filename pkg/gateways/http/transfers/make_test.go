package transfers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/domain/transfers"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
)

func TestMake(t *testing.T) {

	t.Run("should return 201 and a account when it's been sucessfully created", func(t *testing.T) {

		transfer := entities.Transfer{AccountOriginID: 1, AccountDestinationID: 0, Amount: 10}

		useCase := transfers.UseCaseMock{Transfer: transfer}
		h := NewHandler(&useCase)

		request, _ := json.Marshal(transfer)

		newRequest, _ := http.NewRequest(http.MethodPost, "/transfers", bytes.NewReader(request))
		newResponse := httptest.NewRecorder()

		h.Make(newResponse, newRequest)

		createAt := transfer.CreatedAt
		ExpectedCreateAt := createAt.Format(server_http.DateLayout)

		var response Response

		_ = json.Unmarshal(newResponse.Body.Bytes(), &response)

		if newResponse.Code != http.StatusCreated {
			t.Errorf("expected '%d' but got '%d'", http.StatusCreated, newResponse.Code)
		}

		headerType := newResponse.Header().Get("content-type")
		if headerType != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, headerType)
		}

		if response.AccountOriginID != transfer.AccountOriginID {
			t.Errorf("expected '%d' but got '%d'", transfer.AccountOriginID, response.AccountOriginID)
		}

		if response.AccountDestinationID != transfer.AccountDestinationID {
			t.Errorf("expected '%d' but got '%d'", transfer.AccountDestinationID, response.AccountDestinationID)
		}

		if response.Amount != transfer.Amount {
			t.Errorf("expected '%d' but got '%d'", transfer.Amount, response.Amount)
		}

		if response.CreatedAt != ExpectedCreateAt {
			t.Errorf("expected '%s' but got '%s'", ExpectedCreateAt, response.CreatedAt)
		}

	})

	/* 	t.Run("should return 400 and a error message when it failed to decode the request successfully", func(t 		createRequest := Request{AccountOriginID: 1, *testing.T) {

	   	AccountDestinationID: 0, Amount: 10}

	   			storage := transfers_storage.NewStorage()
	   			useCase := transfers_usecase.NewTransferUseCase(storage)
	   			h := NewHandler(useCase)

	   		})  */

}
