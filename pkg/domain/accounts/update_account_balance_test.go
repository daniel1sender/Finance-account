package accounts

import (
	"errors"
	"testing"

	"exemplo.com/pkg/domain/entities"
	"exemplo.com/pkg/store/accounts"
)

func TestAccountUseCase_UpdateAccountBalance(t *testing.T) {

	t.Run("Should return nil when account was updated", func(t *testing.T) {

		storage := accounts.NewAccountStorage()
		AccountUseCase := NewAccountUseCase(storage)
		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("Expected nil error to create a new account but got %s", err)
		}

		storage.UpdateStorage(account.ID, account)

		UpdateAccountError := AccountUseCase.UpdateAccountBalance(account.ID, 20.0)

		if UpdateAccountError != nil {
			t.Errorf("Expected nil but got %s", UpdateAccountError)
		}

	})

	t.Run("Should return an error massage when account don't exists", func(t *testing.T) {

		storage := accounts.NewAccountStorage()
		AccountUseCase := NewAccountUseCase(storage)

		//passando qualquer id, sem criar a conta
		err := AccountUseCase.UpdateAccountBalance("1", 20.0)

		if err != accounts.ErrIDNotFound {
			t.Errorf("Expected %s but got %s", accounts.ErrIDNotFound, err)
		}

	})

	t.Run("Should return an error message when balance account is less than zero", func(t *testing.T) {

		storage := accounts.NewAccountStorage()
		AccountUseCase := NewAccountUseCase(storage)

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("Expected nil error to create a new account but got %s", err)
		}

		storage.UpdateStorage(account.ID, account)

		err = AccountUseCase.UpdateAccountBalance(account.ID, -10)

		if !errors.Is(err, ErrBalanceLessOrEqualZero) {
			t.Errorf("Expected %s but got %s", ErrBalanceLessOrEqualZero, err)
		}

	})
}
