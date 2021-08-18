package store

import "exemplo.com/pkg/domain/entities"

func (s AccountStorage) FindByID(id string) (entities.Account, error) {
	account, ok := s.storage[id]
	if !ok {
		return entities.Account{}, ErrIDNotFound
	}
	return account, nil
}
