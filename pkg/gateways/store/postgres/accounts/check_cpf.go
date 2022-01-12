package accounts

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
)

func (ar AccountRepository) CheckCPF(cpf string) error {
	var CPFaccount string
	if err := ar.QueryRow(context.Background(), "SELECT CPF FROM ACCOUNTS WHERE CPF=$1", cpf).Scan(&CPFaccount); err != nil {
		return err
	} else {
		return accounts.ErrExistingCPF
	}
}
