package accounts

import (
	"context"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (ar AccountRepository) GetAll() ([]entities.Account, error) {
	var users []entities.Account
	var user entities.Account

	rows, err := ar.Query(context.Background(), "SELECT id, name, cpf, secret, balance, created_at FROM accounts")
	if err != nil {
		return []entities.Account{}, fmt.Errorf("error to read table rows: %w", err)
	}

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.CPF, &user.Secret, &user.Balance, &user.CreatedAt)
		if err != nil {
			return []entities.Account{}, err
		}
		users = append(users, user)
	}
	if len(users) == 0{
		return []entities.Account{}, fmt.Errorf("error while listing accounts: %w", accounts.ErrEmptyList)
	}
	return users, nil
}
