package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"

	accounts_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/accounts/usecases"
	transfers_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/transfers"
	accounts_handler "github.com/daniel1sender/Desafio-API/pkg/gateways/http/accounts"
	transfers_handler "github.com/daniel1sender/Desafio-API/pkg/gateways/http/transfers"
	accounts_storage "github.com/daniel1sender/Desafio-API/pkg/gateways/store/memory/accounts"
	transfers_storage "github.com/daniel1sender/Desafio-API/pkg/gateways/store/memory/transfers"
	postgres "github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/accounts"
)

const DatabaseURL = "postgres://postgres:1234@localhost:5432/desafio"

func main() {

	err := postgres.RunMigrations(DatabaseURL)
	if err != nil {
		log.Fatalf("error to run migrations: %v", err)
	}

	dbPool, err := pgxpool.Connect(context.Background(), DatabaseURL)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	defer dbPool.Close()

	accountRepository := accounts.NewStorage(dbPool)
	accountUseCase := accounts_usecase.NewUseCase(accountRepository)
	accountHandler := accounts_handler.NewHandler(accountUseCase)

	accountsMemoryRepository := accounts_storage.NewStorage()
	transferStorage := transfers_storage.NewStorage()
	transferUseCase := transfers_usecase.NewUseCase(transferStorage, accountsMemoryRepository)
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
