package accounts

import "github.com/daniel1sender/Desafio-API/pkg/domain/entities"

func (s AccountStorage) Get() []entities.Account {
	var accountsList []entities.Account
	for _, value := range s.storage {
		accountsList = append(accountsList, value)
	}
	return accountsList
}
