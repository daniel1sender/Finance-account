package accounts

import "github.com/daniel1sender/Desafio-API/pkg/domain/entities"

func (ar AccountRepository) GetAll() []entities.Account {
	var accountsList []entities.Account
	for i, value := range ar.Users {
		account := entities.Account{ID: i, Name: value.Name, CPF: value.CPF, Balance: value.Balance, CreatedAt: value.CreatedAt}
		accountsList = append(accountsList, account)
	}
	return accountsList
}
