package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	accounts_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	transfers_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/transfers"
	accounts_handler "github.com/daniel1sender/Desafio-API/pkg/gateways/http/accounts"
	transfers_handler "github.com/daniel1sender/Desafio-API/pkg/gateways/http/transfers"
	accounts_storage "github.com/daniel1sender/Desafio-API/pkg/gateways/store/accounts"
	transfers_storage "github.com/daniel1sender/Desafio-API/pkg/gateways/store/transfers"
)

func main() {

	transferStorage := transfers_storage.NewStorage()
	transferUseCase := transfers_usecase.NewTransferUseCase(transferStorage)
	transferHandler := transfers_handler.NewHandler(transferUseCase)

	accountStorage := accounts_storage.NewStorage()
	accountUseCase := accounts_usecase.NewUseCase(accountStorage)
	accountHandler := accounts_handler.NewHandler(accountUseCase)

	r := mux.NewRouter()
	r.HandleFunc("/accounts", accountHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/accounts", accountHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/accounts/{id}/balance", accountHandler.GetBalanceByID).Methods(http.MethodGet)
	r.HandleFunc("/transfers", transferHandler.Make).Methods(http.MethodPost)

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("failed to listen and serve: %s", err)
	}
}
