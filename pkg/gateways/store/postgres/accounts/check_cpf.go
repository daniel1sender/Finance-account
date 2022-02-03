package accounts

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/jackc/pgx/v4"
)

const checkCPFStatement = `SELECT 
	cpf 
	FROM accounts 
	WHERE cpf = $1`

func (ar AccountRepository) CheckCPF(ctx context.Context, cpf string) error {
	var CPFaccount string
	err := ar.QueryRow(ctx, checkCPFStatement, cpf).Scan(&CPFaccount)
	if err == pgx.ErrNoRows {
		return nil
	} else if err != nil {
		return err
	}
	return accounts.ErrExistingCPF
}
