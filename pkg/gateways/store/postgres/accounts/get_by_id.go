package accounts

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

const GetByIDStatement = "SELECT id, name, cpf, secret, balance, created_at FROM accounts WHERE id = $1"

func (ar AccountRepository) GetByID(id string) (entities.Account, error) {
	account := entities.Account{}
	err := ar.QueryRow(context.Background(), GetByIDStatement, id).Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
	if err != nil {
		return entities.Account{}, err
	}

	return account, nil
}
