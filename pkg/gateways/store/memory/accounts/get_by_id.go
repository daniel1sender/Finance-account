package accounts

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store"
)

func (s AccountStorage) GetByID(id string) (entities.Account, error) {
	account, ok := s.storage[id]
	if !ok {
		return entities.Account{}, store.ErrIDNotFound
	}
	return account, nil
}
