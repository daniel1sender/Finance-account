package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	accounts_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/accounts/usecases"
	transfers_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/transfers"
	accounts_handler "github.com/daniel1sender/Desafio-API/pkg/gateways/http/accounts"
	transfers_handler "github.com/daniel1sender/Desafio-API/pkg/gateways/http/transfers"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/transfers"
)

func main() {

	accountRepository := accounts.NewStorage()
	accountUseCase := accounts_usecase.NewUseCase(accountRepository)
	accountHandler := accounts_handler.NewHandler(accountUseCase)

	transferStorage := transfers.NewStorage()
	transferUseCase := transfers_usecase.NewUseCase(transferStorage, accountRepository)
	transferHandler := transfers_handler.NewHandler(transferUseCase)

	r := mux.NewRouter()
	r.HandleFunc("/accounts", accountHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/accounts", accountHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/accounts/{id}/balance", accountHandler.GetBalanceByID).Methods(http.MethodGet)
	r.HandleFunc("/transfers", transferHandler.Make).Methods(http.MethodPost)

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("failed to listen and serve: %s", err)
	}
}
