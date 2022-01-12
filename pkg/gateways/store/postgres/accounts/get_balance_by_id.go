package accounts

import (
	"context"
	"log"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (ar AccountRepository) GetBalanceByID(id string) (int, error) {
	account := entities.Account{}
	err := ar.QueryRow(context.Background(), "SELECT BALANCE FROM ACCOUNTS WHERE ID = $1", id).Scan(&account.Balance)
	if err != nil {
		log.Print(err)
	}

	return account.Balance, nil
}
