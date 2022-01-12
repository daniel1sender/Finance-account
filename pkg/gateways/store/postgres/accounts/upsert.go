package accounts

import (
	"context"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (ar AccountRepository) Upsert(account entities.Account) error {
	if _, err := ar.Exec(context.Background(), "INSERT INTO ACCOUNTS(id, name, cpf, secret, balance, createdat) VALUES ($1, $2, $3, $4, $5, $6)", account.ID, account.Name, account.CPF, account.Secret, account.Balance, account.CreatedAt); err != nil {
		return fmt.Errorf("unable to insert the account due to: %v", err)
	}
	return nil
}
