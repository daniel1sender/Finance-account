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
	accounts_storage "github.com/daniel1sender/Desafio-API/pkg/gateways/store/accounts"
)

func TestMake(t *testing.T) {

	t.Run("should return 201 and a transfer when it's been sucessfully created", func(t *testing.T) {

		transfer := entities.Transfer{AccountOriginID: "1", AccountDestinationID: "0", Amount: 10}
		useCase := transfers.UseCaseMock{Transfer: transfer}
		h := NewHandler(&useCase)

		createRequest := Request{transfer.AccountOriginID, transfer.AccountDestinationID, transfer.Amount}
		request, _ := json.Marshal(createRequest)
		newRequest, _ := http.NewRequest(http.MethodPost, "/transfers", bytes.NewReader(request))
		newResponse := httptest.NewRecorder()
		h.Make(newResponse, newRequest)

		ExpectedCreateAt := transfer.CreatedAt.Format(server_http.DateLayout)

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
			t.Errorf("expected '%s' but got '%s'", transfer.AccountOriginID, response.AccountOriginID)
		}

		if response.AccountDestinationID != transfer.AccountDestinationID {
			t.Errorf("expected '%s' but got '%s'", transfer.AccountDestinationID, response.AccountDestinationID)
		}

		if response.Amount != transfer.Amount {
			t.Errorf("expected '%d' but got '%d'", transfer.Amount, response.Amount)
		}

		if response.CreatedAt != ExpectedCreateAt {
			t.Errorf("expected '%s' but got '%s'", ExpectedCreateAt, response.CreatedAt)
		}

	})

	t.Run("should return 400 and a error message when it failed to decode the request", func(t *testing.T) {

		useCase := transfers.UseCaseMock{}
		h := NewHandler(&useCase)
		b := []byte{}
		newRequest, _ := http.NewRequest(http.MethodPost, "transfers", bytes.NewBuffer(b))
		newResponse := httptest.NewRecorder()
		h.Make(newResponse, newRequest)

		var responseReason server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected status '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}

		expected := "invalid request body"
		if responseReason.Reason != expected {
			t.Errorf("expected '%s' but got '%s'", expected, responseReason.Reason)
		}

	})

	t.Run("should return 400 and a error message when amount is less or equal zero", func(t *testing.T) {

		transfer := entities.Transfer{AccountOriginID: "1", AccountDestinationID: "0", Amount: -10}
		useCase := transfers.UseCaseMock{Transfer: transfer, Error: entities.ErrAmountLessOrEqualZero}
		h := NewHandler(&useCase)

		createRequest := Request{transfer.AccountOriginID, transfer.AccountDestinationID, transfer.Amount}
		request, _ := json.Marshal(createRequest)
		newRequest, _ := http.NewRequest(http.MethodPost, "/transfers", bytes.NewBuffer(request))
		newResponse := httptest.NewRecorder()
		h.Make(newResponse, newRequest)

		var responseReason server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected status '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}

		if responseReason.Reason != entities.ErrAmountLessOrEqualZero.Error() {
			t.Errorf("expected '%s' but got '%s'", entities.ErrAmountLessOrEqualZero.Error(), responseReason.Reason)
		}

	})

	t.Run("should return 400 and a error message when the transfer is to the same account", func(t *testing.T) {

		transfer := entities.Transfer{AccountOriginID: "0", AccountDestinationID: "0", Amount: 10}
		useCase := transfers.UseCaseMock{Transfer: transfer, Error: entities.ErrSameAccountTransfer}
		h := NewHandler(&useCase)

		createRequest := Request{transfer.AccountOriginID, transfer.AccountDestinationID, transfer.Amount}
		request, _ := json.Marshal(createRequest)
		newRequest, _ := http.NewRequest(http.MethodPost, "/transfers", bytes.NewBuffer(request))
		newResponse := httptest.NewRecorder()
		h.Make(newResponse, newRequest)

		var responseReason server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected status '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}

		if responseReason.Reason != entities.ErrSameAccountTransfer.Error() {
			t.Errorf("expected '%s' but got '%s'", entities.ErrSameAccountTransfer.Error(), responseReason.Reason)
		}
	})

	t.Run("should return 400 and an error message when ids of transfer isn't found", func(t *testing.T) {

		transfer := entities.Transfer{AccountOriginID: "0", AccountDestinationID: "1", Amount: 10}
		useCase := transfers.UseCaseMock{Transfer: transfer, Error: accounts_storage.ErrIDNotFound}
		h := NewHandler(&useCase)

		createRequest := Request{transfer.AccountOriginID, transfer.AccountDestinationID, transfer.Amount}
		request, _ := json.Marshal(createRequest)
		newRequest, _ := http.NewRequest(http.MethodPost, "/transfers", bytes.NewBuffer(request))
		newResponse := httptest.NewRecorder()
		h.Make(newResponse, newRequest)

		var responseReason server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected status '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}

		if responseReason.Reason != accounts_storage.ErrIDNotFound.Error() {
			t.Errorf("expected '%s' but got '%s'", accounts_storage.ErrIDNotFound.Error(), responseReason.Reason)
		}
	})

	t.Run("should return 400 and error message when origin account doesn't have sufficient funds", func(t *testing.T) {

		transfer := entities.Transfer{AccountOriginID: "0", AccountDestinationID: "1", Amount: 10}
		useCase := transfers.UseCaseMock{Transfer: transfer, Error: entities.ErrInsufficientFunds}
		h := NewHandler(&useCase)

		createRequest := Request{transfer.AccountOriginID, transfer.AccountDestinationID, transfer.Amount}
		request, _ := json.Marshal(createRequest)
		newRequest, _ := http.NewRequest(http.MethodPost, "/transfers", bytes.NewBuffer(request))
		newResponse := httptest.NewRecorder()
		h.Make(newResponse, newRequest)

		var responseReason server_http.Error
		json.Unmarshal(newResponse.Body.Bytes(), &responseReason)

		if newResponse.Code != http.StatusBadRequest {
			t.Errorf("expected status '%d' but got '%d'", http.StatusBadRequest, newResponse.Code)
		}

		if newResponse.Header().Get("content-type") != server_http.JSONContentType {
			t.Errorf("expected '%s' but got '%s'", server_http.JSONContentType, newResponse.Header().Get("content-type"))
		}

		if responseReason.Reason != entities.ErrInsufficientFunds.Error() {
			t.Errorf("expected '%s' but got '%s'", entities.ErrInsufficientFunds.Error(), responseReason.Reason)
		}

	})

}
