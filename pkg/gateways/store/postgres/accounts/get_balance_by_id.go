package accounts

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/jackc/pgx/v4"
)

func (ar AccountRepository) GetBalanceByID(id string) (int, error) {
	account := entities.Account{}
	err := ar.QueryRow(context.Background(), "SELECT balance FROM accounts WHERE id = $1", id).Scan(&account.Balance)
	if err == pgx.ErrNoRows {
		return 0, accounts.ErrAccountNotFound
	} else if err != nil {
		return 0, err
	}

	return account.Balance, nil
}
