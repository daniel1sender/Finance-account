package accounts

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store"
)

func (ar AccountRepository) GetByID(id string) (entities.Account, error) {
	account, ok := ar.Users[id]
	if !ok {
		return entities.Account{}, store.ErrIDNotFound
	}
	return account, nil
}
