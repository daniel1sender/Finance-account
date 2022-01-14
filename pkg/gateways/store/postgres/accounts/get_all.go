package accounts

import (
	"context"
	"log"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (ar AccountRepository) GetAll() []entities.Account {
	var users []entities.Account
	var user entities.Account

	rows, err := ar.Query(context.Background(), "SELECT id, name, cpf, secret, balance, created_at FROM accounts")
	if err != nil {
		log.Printf("error to read table rows: %v", err)
	}

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.CPF, &user.Secret, &user.Balance, &user.CreatedAt)
		if err != nil {
			log.Print(err)
		}
		users = append(users, user)
	}
	return users
}
