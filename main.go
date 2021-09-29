package main

import (
	"log"
	"net/http"

	accounts_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	accounts_handler "github.com/daniel1sender/Desafio-API/pkg/gateways/http/accounts"
	accounts_storage "github.com/daniel1sender/Desafio-API/pkg/gateways/store/accounts"
	"github.com/gorilla/mux"
)

func main() {

	accountStorage := accounts_storage.NewStorage()
	accountUseCase := accounts_usecase.NewUseCase(accountStorage)
	accountHandler := accounts_handler.NewHandler(accountUseCase)

	r := mux.NewRouter()
	r.HandleFunc("/accounts", accountHandler.Create).Methods(http.MethodPost) // accountHandler.Create()
	r.HandleFunc("/accounts", accountHandler.Get).Methods(http.MethodGet)
	r.HandleFunc("/accounts/{id}/balance", accountHandler.GetBalanceByID).Methods(http.MethodGet)

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("failed to listen and serve: %s", err)
	}
}
