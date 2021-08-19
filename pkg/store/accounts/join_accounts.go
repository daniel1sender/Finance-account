package accounts

import "exemplo.com/pkg/domain/entities"

func (s AccountStorage) JoinAccounts() []entities.Account {
	var AccountsList []entities.Account
	for _, value := range s.storage {
		AccountsList = append(AccountsList, value)
	}
	return AccountsList
}
