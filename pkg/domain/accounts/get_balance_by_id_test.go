package accounts

import (
	"testing"

	"exemplo.com/pkg/domain/entities"
	"exemplo.com/pkg/store/accounts"
)

func TestAccountUseCase_GetBalanceByID(t *testing.T) {

	t.Run("should return an account when id is found", func(t *testing.T) {

		storage := accounts.NewAccountStorage()
		AccountUseCase := NewAccountUseCase(storage)
		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("Expected nil error to create a new account but got '%s'", err)
		}
		storage.UpdateStorage(account.ID, account)

		getBalance, err := AccountUseCase.GetBalanceByID(account.ID)

		if getBalance == 0 {
			t.Error("Expected balance account different from 0")
		}

		if err != nil {
			t.Errorf("Expected nil error but got '%s'", err)
		}

	})

	t.Run("should return a blank account when id isn't found", func(t *testing.T) {

		storage := accounts.NewAccountStorage()
		AccountUseCase := NewAccountUseCase(storage)
		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("Expected nil error to create a new account but got '%s'", err)
		}

		getBalance, err := AccountUseCase.GetBalanceByID(account.ID)

		if getBalance != 0 {
			t.Error("Expected account balance equal zero")
		}

		if err == nil {
			t.Errorf("Expected nil error but got '%s'", err)
		}

	})

}