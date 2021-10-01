package accounts

import (
	"errors"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

var (
	ErrExistingCPF = errors.New("cpf informed is invalid")
)

func (au AccountUseCase) Create(name, cpf, secret string, balance int) (entities.Account, error) {

	if err := au.storage.CheckCPF(cpf); err != nil {
		return entities.Account{}, ErrExistingCPF
	}

	account, err := entities.NewAccount(name, cpf, secret, balance)
	if err != nil {
		return entities.Account{}, fmt.Errorf("error while creating an account: %w", err)
	}

	au.storage.Upsert(account.ID, account)

	return account, nil
}
