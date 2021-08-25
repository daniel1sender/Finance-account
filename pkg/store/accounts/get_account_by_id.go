package accounts

import "github.com/daniel1sender/Desafio-API/pkg/domain/entities"

func (s AccountStorage) GetAccountByID(id string) (entities.Account, error) {
	account, ok := s.storage[id]
	if !ok {
		return entities.Account{}, ErrIDNotFound
	}
	return account, nil
}
