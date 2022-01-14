package accounts

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (ar AccountRepository) GetByID(id string) (entities.Account, error) {
	account := entities.Account{}
	err := ar.QueryRow(context.Background(), "SELECT id, name, cpf, secret, balance, created_at FROM accounts WHERE id = $1", id).Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
	if err != nil {
		return entities.Account{}, err
	}

	return account, nil
}
