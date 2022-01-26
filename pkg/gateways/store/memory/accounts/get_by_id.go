package accounts

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (s AccountStorage) GetByID(id string) (entities.Account, error) {
	account, ok := s.storage[id]
	if !ok {
		return entities.Account{}, accounts.ErrAccountNotFound
	}
	return account, nil
}
