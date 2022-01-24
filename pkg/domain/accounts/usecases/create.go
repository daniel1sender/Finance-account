package usecases

import (
	"errors"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/jackc/pgx/v4"
)

func (au AccountUseCase) Create(name, cpf, secret string, balance int) (entities.Account, error) {

	err := au.storage.CheckCPF(cpf)
	if err != nil {
		switch {
		case errors.Is(err, accounts.ErrExistingCPF):
			return entities.Account{}, fmt.Errorf("%w", accounts.ErrExistingCPF)
		case !errors.Is(err, pgx.ErrNoRows):
			return entities.Account{}, fmt.Errorf("%w", err)
		}
	}

	account, err := entities.NewAccount(name, cpf, secret, balance)
	if err != nil {
		return entities.Account{}, fmt.Errorf("error while creating an account: %w", err)
	}

	err = au.storage.Upsert(account)
	if err != nil {
		return entities.Account{}, fmt.Errorf("%w", err)
	}

	return account, nil
}
