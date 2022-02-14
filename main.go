package main

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"

	accounts_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/accounts/usecases"
	transfers_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/transfers"
	accounts_handler "github.com/daniel1sender/Desafio-API/pkg/gateways/http/accounts"
	transfers_handler "github.com/daniel1sender/Desafio-API/pkg/gateways/http/transfers"
	postgres "github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/transfers"
)
type Config struct {
	DatabaseURL string `envconfig:"DB_URL"`
	Port        string `envconfig:"API_PORT"`
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	entry := logrus.NewEntry(log)
	var s Config
	err := envconfig.Process("", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = postgres.RunMigrations(s.DatabaseURL)
	if err != nil {
		log.Fatalf("error to run migrations: %v", err)
	}

	dbPool, err := pgxpool.Connect(context.Background(), s.DatabaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	defer dbPool.Close()

	accountRepository := accounts.NewStorage(dbPool)
	accountUseCase := accounts_usecase.NewUseCase(accountRepository)
	accountHandler := accounts_handler.NewHandler(accountUseCase, entry)

	transferStorage := transfers.NewStorage(dbPool)
	transferUseCase := transfers_usecase.NewUseCase(transferStorage, accountRepository)
	transferHandler := transfers_handler.NewHandler(transferUseCase, entry)

	r := mux.NewRouter()
	r.HandleFunc("/accounts", accountHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/accounts", accountHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/accounts/{id}/balance", accountHandler.GetBalanceByID).Methods(http.MethodGet)
	r.HandleFunc("/transfers", transferHandler.Make).Methods(http.MethodPost)

	if err := http.ListenAndServe(s.Port, r); err != nil {
		log.Fatalf("failed to listen and serve: %s", err)
	}
}
