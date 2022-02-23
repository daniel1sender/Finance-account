package accounts

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/jackc/pgx/v4"
)

const getByCPFStatement = "SELECT id, name, cpf, secret, balance, created_at FROM accounts WHERE cpf = $1"

func (ar AccountRepository) GetByCPF(ctx context.Context, cpf string) (entities.Account, error) {
	account := entities.Account{}
	err := ar.QueryRow(ctx, getByCPFStatement, cpf).Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
	if err == pgx.ErrNoRows {
		return entities.Account{}, accounts.ErrAccountNotFound
	} else if err != nil {
		return entities.Account{}, err
	}

	return account, nil
}
