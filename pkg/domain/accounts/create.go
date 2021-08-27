package accounts

import (
	"errors"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

var (
	ErrCreatingNewAccount = errors.New("error while creating an account")
	ErrExistingCPF        = errors.New("cpf informed is invalid")
)

func (au AccountUseCase) Create(name, cpf, secret string, balance int) (entities.Account, error) {

	if err := au.storage.CheckCPF(cpf); err != nil {
		return entities.Account{}, ErrExistingCPF
	}

	account, err := entities.NewAccount(name, cpf, secret, balance)
	if err != nil {
		return entities.Account{}, ErrCreatingNewAccount
	}

	au.storage.Upsert(account.ID, account)

	return account, nil
}
