package main

import (
	"context"
	"net/http"
	"os"

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
	DatabaseURL string `envconfig:"DB_URL" required:"true"`
	Port        string `envconfig:"API_PORT" required:"true" default:":3000"`
}

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	log := logrus.NewEntry(logger)

	keysVariables := []string{"DB_URL", "API_PORT"}
	valuesVariables := []string{"postgres://postgres:1234@localhost:5432/desafio", ":3000"}
	for i := 0; i < len(keysVariables); i++ {
		err := os.Setenv(keysVariables[i], valuesVariables[i])
		if err != nil {
			log.WithError(err).Fatal("error while setting environment variables")
		}
	}

	var apiConfig Config
	err := envconfig.Process("", &apiConfig)
	if err != nil {
		log.WithError(err).Fatal("error while processing environment variables")
	}

	err = postgres.RunMigrations(apiConfig.DatabaseURL)
	if err != nil {
		log.WithError(err).Fatal("error while running migrations")
	}

	dbPool, err := pgxpool.Connect(context.Background(), apiConfig.DatabaseURL)
	if err != nil {
		log.WithError(err).Fatal("error while connecting with the database")
	}

	defer dbPool.Close()

	accountRepository := accounts.NewStorage(dbPool)
	accountUseCase := accounts_usecase.NewUseCase(accountRepository)
	accountHandler := accounts_handler.NewHandler(accountUseCase, log)

	transferStorage := transfers.NewStorage(dbPool)
	transferUseCase := transfers_usecase.NewUseCase(transferStorage, accountRepository)
	transferHandler := transfers_handler.NewHandler(transferUseCase, log)

	r := mux.NewRouter()
	r.HandleFunc("/accounts", accountHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/accounts", accountHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/accounts/{id}/balance", accountHandler.GetBalanceByID).Methods(http.MethodGet)
	r.HandleFunc("/transfers", transferHandler.Make).Methods(http.MethodPost)

	if err := http.ListenAndServe(apiConfig.Port, r); err != nil {
		log.WithError(err).Fatal("failed to listen and serve")
	}
}
