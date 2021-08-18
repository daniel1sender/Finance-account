package accounts

import (
	"errors"
	"testing"

	"exemplo.com/pkg/domain/entities"
	"exemplo.com/pkg/store"
)

func TestAccountUseCase_CreateAccount(t *testing.T) {
	t.Run("should successfully create an account and return it", func(t *testing.T) {

		storage := store.NewAccountStorage()
		accountUsecase := NewAccountUseCase(storage)
		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		createdAccount, err := accountUsecase.CreateAccount(name, cpf, secret, balance)

		if err != nil {
			t.Error("Expected nil error but got %w", err)
		}

		if createdAccount == (entities.Account{}) {
			t.Errorf("Expected an account but got %v", createdAccount)
		}

	})

	t.Run("should return error when trying to create account with already created cpf account", func(t *testing.T) {

		storage := store.NewAccountStorage()
		accountUsecase := NewAccountUseCase(storage)

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		createdAccount, err := accountUsecase.CreateAccount(name, cpf, secret, balance)

		if err != nil {
			t.Errorf("Expected nil error but got %s", err)
		}

		if createdAccount == (entities.Account{}) {
			t.Errorf("Expected %+v but got %+v", entities.Account{}, createdAccount)
		}

		createdAccount1, err1 := accountUsecase.CreateAccount(name, cpf, secret, balance)

		if !errors.Is(err1, store.ErrExistingCPF) {
			t.Errorf("Expected %s but got %s", store.ErrExistingCPF, err1)
		}

		if createdAccount1 != (entities.Account{}) {
			t.Errorf("Expected %+v but got %+v", entities.Account{}, createdAccount1)
		}

	})
}
