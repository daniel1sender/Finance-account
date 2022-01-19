package accounts

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (ar AccountRepository) GetBalanceByID(id string) (int, error) {
	account := entities.Account{}
	err := ar.QueryRow(context.Background(), "SELECT balance FROM accounts WHERE id = $1", id).Scan(&account.Balance)
	if err != nil {
		return 0, accounts.ErrIDNotFound
	}

	return account.Balance, nil
}
