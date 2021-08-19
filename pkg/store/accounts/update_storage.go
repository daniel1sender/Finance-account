package accounts

import "exemplo.com/pkg/domain/entities"

func (s AccountStorage) UpdateStorage(id string, account entities.Account) {
	s.storage[id] = account
}
