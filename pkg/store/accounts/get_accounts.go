package accounts

import "github.com/daniel1sender/Desafio-API/pkg/domain/entities"

func (s AccountStorage) GetAccounts() []entities.Account {
	var AccountsList []entities.Account
	for _, value := range s.storage {
		AccountsList = append(AccountsList, value)
	}
	return AccountsList
}
