package accounts

import (
	"context"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

const upsertStatement = `INSERT INTO accounts(
	id, 
	name, 
	cpf, 
	secret, 
	balance, 
	created_at
	) VALUES (
		$1, 
		$2, 
		$3, 
		$4, 
		$5, 
		$6
	) ON CONFLICT (id) DO UPDATE SET balance = 
	EXCLUDED.balance`

func (ar AccountRepository) Upsert(account entities.Account) error {
	if _, err := ar.Exec(context.Background(), upsertStatement, account.ID, account.Name, account.CPF, account.Secret, account.Balance, account.CreatedAt); err != nil {
		return fmt.Errorf("unable to insert the account due to: %v", err)
	}
	return nil
}
