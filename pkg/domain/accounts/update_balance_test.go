package accounts

import (
	"errors"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/accounts"
)

func TestAccountUseCase_UpdateBalance(t *testing.T) {

	t.Run("should return an account and null error when account was updated", func(t *testing.T) {

		storage := accounts.NewStorage()
		accountUseCase := NewUseCase(storage)
		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}

		storage.Upsert(account.ID, account)

		updateAccountError := accountUseCase.UpdateBalance(account.ID, 20.0)

		if updateAccountError != nil {
			t.Errorf("expected no error but got '%s'", updateAccountError)
		}

	})

	t.Run("should return an empty account an error when account don't exists", func(t *testing.T) {

		storage := accounts.NewStorage()
		accountUseCase := NewUseCase(storage)

		//passando qualquer id, sem criar a conta
		err := accountUseCase.UpdateBalance("1", 20.0)

		if err != accounts.ErrIDNotFound {
			t.Errorf("expected '%s' but got '%s'", accounts.ErrIDNotFound, err)
		}

	})

	t.Run("should return an empty account and an error when balance account is less than zero", func(t *testing.T) {

		storage := accounts.NewStorage()
		accountUseCase := NewUseCase(storage)

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}

		storage.Upsert(account.ID, account)

		err = accountUseCase.UpdateBalance(account.ID, -10)

		if !errors.Is(err, ErrBalanceLessZero) {
			t.Errorf("expected '%s' but got '%s'", ErrBalanceLessZero, err)
		}

	})
}
