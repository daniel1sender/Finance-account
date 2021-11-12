package accounts

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store"
)

func (ar accountRepository) GetByID(id string) (entities.Account, error) {
	account, ok := ar.users[id]
	if !ok {
		return entities.Account{}, store.ErrIDNotFound
	}
	return account, nil
}
