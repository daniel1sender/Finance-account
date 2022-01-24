package accounts

import (
	"context"
	"errors"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/jackc/pgx/v4"
)

func (ar AccountRepository) CheckCPF(cpf string) error {
	var CPFaccount string
	err := ar.QueryRow(context.Background(), "SELECT cpf FROM accounts WHERE cpf = $1", cpf).Scan(&CPFaccount)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil
		default:
			return err
		}

	}
	return accounts.ErrExistingCPF
}
