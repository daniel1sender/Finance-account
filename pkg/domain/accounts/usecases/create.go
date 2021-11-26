package usecases

import (
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (au AccountUseCase) Create(name, cpf, secret string, balance int) (entities.Account, error) {

	if err := au.storage.CheckCPF(cpf); err != nil {
		return entities.Account{}, accounts.ErrExistingCPF
	}

	account, err := entities.NewAccount(name, cpf, secret, balance)
	if err != nil {
		return entities.Account{}, fmt.Errorf("error while creating an account: %w", err)
	}

	au.storage.Upsert(account)

	return account, nil
}