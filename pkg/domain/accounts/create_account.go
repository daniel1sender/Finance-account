package accounts

import (
	"errors"

	"exemplo.com/pkg/domain/entities"
)

var (
	ErrToCallNewAccount = errors.New("error to call function new account")
)

func (au AccountUseCase) CreateAccount(name, cpf, secret string, balance int) (entities.Account, error) {

	if err := au.storage.GetCPF(cpf); err != nil {
		return entities.Account{}, err
	}

	account, err := entities.NewAccount(name, cpf, secret, balance)
	if err != nil {
		return entities.Account{}, ErrToCallNewAccount
	}

	au.storage.UpdateStorage(account.ID, account)

	return account, nil
}
