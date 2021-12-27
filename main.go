package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts/usecases"
	transfers_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/transfers"
	accounts_handler "github.com/daniel1sender/Desafio-API/pkg/gateways/http/accounts"
	transfers_handler "github.com/daniel1sender/Desafio-API/pkg/gateways/http/transfers"
	accounts_repository "github.com/daniel1sender/Desafio-API/pkg/gateways/store/files/accounts"
	transfers_repository "github.com/daniel1sender/Desafio-API/pkg/gateways/store/files/transfers"
)

func main() {

	transferFile := "Transfer_Respository.json"
	openTransferFile, err := os.OpenFile(transferFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error to open file: %v", err)
	}

	transferRepository := transfers_repository.NewStorage(openTransferFile)
	transferUseCase := transfers_usecase.NewUseCase(transferRepository)
	transferHandler := transfers_handler.NewHandler(transferUseCase)

	accountFile := "Account_Repository.json"
	openAccountFile, err := os.OpenFile(accountFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error to open file: %v", err)
	}

	accountRepository := accounts_repository.NewStorage(openAccountFile)
	accountUseCase := usecases.NewUseCase(accountRepository)
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
