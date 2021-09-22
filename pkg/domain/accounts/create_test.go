package accounts

import (
	"errors"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/accounts"
)

func TestAccountUseCase_Create(t *testing.T) {
	t.Run("should successfully create an account and return it", func(t *testing.T) {

		storage := accounts.NewStorage()
		accountUsecase := NewUseCase(storage)
		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		createdAccount, err := accountUsecase.Create(name, cpf, secret, balance)

		if err != nil {
			t.Errorf("expected no error but got '%s'", err)
		}

		if createdAccount == (entities.Account{}) {
			t.Errorf("expected an account but got %+v", createdAccount)
		}

	})

	t.Run("should return error when trying to create account with already created cpf account", func(t *testing.T) {

		storage := accounts.NewStorage()
		accountUsecase := NewUseCase(storage)

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		createdAccount, err := accountUsecase.Create(name, cpf, secret, balance)

		if err != nil {
			t.Errorf("expected no error but got '%s'", err)
		}

		if createdAccount == (entities.Account{}) {
			t.Errorf("expected %+v but got %+v", entities.Account{}, createdAccount)
		}

		createdAccount1, err1 := accountUsecase.Create(name, cpf, secret, balance)

		if !errors.Is(err1, ErrExistingCPF) {
			t.Errorf("expected '%s' but got '%s'", accounts.ErrExistingCPF, err1)
		}

		if createdAccount1 != (entities.Account{}) {
			t.Errorf("expected %+v but got %+v", entities.Account{}, createdAccount1)
		}

	})
}