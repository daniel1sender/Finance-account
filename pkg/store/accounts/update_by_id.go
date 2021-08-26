package accounts

import "github.com/daniel1sender/Desafio-API/pkg/domain/entities"

func (s AccountStorage) Upsert(id string, account entities.Account) {
	s.storage[id] = account
}
