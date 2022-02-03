package accounts

import (
	"context"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/jackc/pgx/v4"
)

const getAllStatement = `SELECT 
	id,
	name,
	cpf,
	secret,
	balance,
	created_at FROM accounts`

func (ar AccountRepository) GetAll(ctx context.Context) ([]entities.Account, error) {
	var users []entities.Account
	var user entities.Account

	rows, err := ar.Query(ctx, getAllStatement)
	if err == pgx.ErrNoRows {
		return []entities.Account{}, accounts.ErrAccountNotFound
	} else if err != nil {
		return []entities.Account{}, err
	}

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.CPF, &user.Secret, &user.Balance, &user.CreatedAt)
		if err != nil {
			return []entities.Account{}, err
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		return []entities.Account{}, fmt.Errorf("error while listing accounts: %w", accounts.ErrEmptyList)
	}
	return users, nil
}
