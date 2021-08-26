package accounts

import (
	"errors"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

var (
	ErrToCreateNewAccount = errors.New("error to create a new account")
	ErrExistingCPF        = errors.New("cpf informed is invalid")
)

func (au AccountUseCase) CreateAccount(name, cpf, secret string, balance int) (entities.Account, error) {

	if err := au.storage.CheckCPF(cpf); err != nil {
		return entities.Account{}, ErrExistingCPF
	}

	account, err := entities.NewAccount(name, cpf, secret, balance)
	if err != nil {
		return entities.Account{}, ErrToCreateNewAccount
	}

	au.storage.Upsert(account.ID, account)

	return account, nil
}
