package accounts

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (ar AccountRepository) GetByID(id string) (entities.Account, error) {
	account := entities.Account{}
	err := ar.QueryRow(context.Background(), "SELECT ID, NAME, CPF, SECRET, BALANCE, CREATEDAT FROM ACCOUNTS", id).Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
	if err != nil {
		return entities.Account{}, err
	}

	return account, nil
}
