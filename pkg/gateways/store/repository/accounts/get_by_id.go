package accounts

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (ar accountRepository) GetByID(id string) (entities.Account, error) {
	account, ok := ar.users[id]
	if !ok {
		return entities.Account{}, accounts.ErrIDNotFound
	}
	return account, nil
}
