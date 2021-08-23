package accounts

import (
	"testing"

	"exemplo.com/pkg/domain/entities"
	"exemplo.com/pkg/store/accounts"
)

func TestAccountUseCase_CheckAccounts(t *testing.T) {

	t.Run("should return nil when accounts have already been created", func(t *testing.T) {

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

		CheckAccountsError := AccountUseCase.CheckAccounts(account.ID)

		if CheckAccountsError != nil {
			t.Error("expected nil when account exists")
		}

	})

	t.Run("should return an error message when id isn't found", func(t *testing.T) {

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

		CheckAccountsError := AccountUseCase.CheckAccounts(account.ID)

		if CheckAccountsError == nil {
			t.Error("expected a error message")
		}

	})

}
