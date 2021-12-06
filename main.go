package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts/usecases"
	transfers_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/transfers"
	accounts_handler "github.com/daniel1sender/Desafio-API/pkg/gateways/http/accounts"
	transfers_handler "github.com/daniel1sender/Desafio-API/pkg/gateways/http/transfers"
	accounts_repository "github.com/daniel1sender/Desafio-API/pkg/gateways/store/files/accounts"
	transfers_repository "github.com/daniel1sender/Desafio-API/pkg/gateways/store/files/transfers"
)

func main() {

	transferRepository := transfers_repository.NewStorage()
	transferUseCase := transfers_usecase.NewUseCase(transferRepository)
	transferHandler := transfers_handler.NewHandler(transferUseCase)

	//transferMemory := transfers_memory.NewStorage()
	//transferUseCase := transfers_usecase.NewUseCase(transferMemory)
	//transferHandler := transfers_handler.NewHandler(transferUseCase)

	//accountMemory := accounts_memory.NewStorage()
	//accountUseCase := accounts_usecase.NewUseCase(accountMemory)
	//accountHandler := accounts_handler.NewHandler(accountUseCase)

	accountRepository := accounts_repository.NewStorage()
	accountUseCase := usecases.NewUseCase(accountRepository)
	accountHandler := accounts_handler.NewHandler(accountUseCase)

	/* 	name := "John Doe"
	   	cpf := "11111111030"
	   	secret := "123"
	   	balance := 10
	   	account, _ := entities.NewAccount(name, cpf, secret, balance)
	   	accountRepository := accounts_repository.NewStorage()
	   	accountRepository.Upsert(account) */

	r := mux.NewRouter()
	r.HandleFunc("/accounts", accountHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/accounts", accountHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/accounts/{id}/balance", accountHandler.GetBalanceByID).Methods(http.MethodGet)
	r.HandleFunc("/transfers", transferHandler.Make).Methods(http.MethodPost)

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("failed to listen and serve: %s", err)
	}
}
