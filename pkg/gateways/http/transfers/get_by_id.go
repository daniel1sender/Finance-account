package transfers

import (
	"encoding/json"
	"net/http"

	"github.com/daniel1sender/Desafio-API/pkg/config"
	server_http "github.com/daniel1sender/Desafio-API/pkg/gateways/http"
	"github.com/gorilla/mux"
)

type Transfer struct {
	Id                   string `json:"id"`
	AccountOriginID      string `json:"account_origin_id"`
	AccountDestinationID string `json:"account_destination_id"`
	Amount               int    `json:"amount"`
	CreatedAt            string `json:"create_at"`
}

type ResponseGet struct {
	List []Transfer `json:"list"`
}

func (h Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	log := h.logger
	accountID := mux.Vars(r)["id"]

	token, err := h.loginUseCase.GetTokenByID(r.Context(), accountID)
	if err != nil {
		w.Header().Add("Content-Type", server_http.JSONContentType)
		response := server_http.Error{Reason: "token not found"}
		log.WithError(err).Error("error while getting token")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	config, err := config.GetConfig()
	if err != nil {
		w.Header().Add("Content-Type", server_http.JSONContentType)
		response := server_http.Error{Reason: "unable to get environment variables"}
		log.WithError(err).Error("unable to get environment variables")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Add("Content-Type", server_http.JSONContentType)
	transfers, err := h.useCase.GetByID(r.Context(), accountID, token, config.TokenSecret)
	if err != nil {
		w.Header().Add("Content-Type", server_http.JSONContentType)
		response := server_http.Error{Reason: "unable to get the list of transfers"}
		log.WithError(err).Error("unable to get the list of transfers")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := ResponseGet{}
	for _, value := range transfers {
		transfer := Transfer{value.ID, value.AccountOriginID, value.AccountDestinationID, value.Amount, value.CreatedAt.Format(server_http.DateLayout)}
		response.List = append(response.List, transfer)
	}

	json.NewEncoder(w).Encode(transfers)

}
