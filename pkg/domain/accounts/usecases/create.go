package usecases

import (
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (au AccountUseCase) Create(name, cpf, secret string, balance int) (entities.Account, error) {

	if err := au.storage.CheckCPF(cpf); err != nil {
		return entities.Account{}, fmt.Errorf("error while checking if the CPF exists: %w", err)
	}

	account, err := entities.NewAccount(name, cpf, secret, balance)
	if err != nil {
		return entities.Account{}, fmt.Errorf("error while creating an account: %w", err)
	}

	err = au.storage.Upsert(account)
	if err != nil {
		return entities.Account{}, fmt.Errorf("error while inserting an account: %w", err)
	}

	return account, nil
}
