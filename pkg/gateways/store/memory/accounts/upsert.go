package accounts

import "github.com/daniel1sender/Desafio-API/pkg/domain/entities"

func (s AccountStorage) Upsert(account entities.Account) error {
	s.storage[account.ID] = account
	return nil
}
