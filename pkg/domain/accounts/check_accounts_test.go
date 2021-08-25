package accounts

import (
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/store/accounts"
)

func TestAccountUseCase_CheckAccounts(t *testing.T) {

	t.Run("should return nil error when accounts have already been created", func(t *testing.T) {

		storage := accounts.NewAccountStorage()
		AccountUseCase := NewAccountUseCase(storage)
		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("expected nil error to create a new account but got '%s'", err)
		}

		storage.UpdateByID(account.ID, account)

		CheckAccountsError := AccountUseCase.CheckAccounts(account.ID)

		if CheckAccountsError != nil {
			t.Errorf("expected nil error but got '%s'", CheckAccountsError)
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
			t.Errorf("expected nil error to create a new account but got '%s'", err)
		}

		CheckAccountsError := AccountUseCase.CheckAccounts(account.ID)

		if CheckAccountsError == nil {
			t.Error("expected a error message but got nil")
		}

	})

}
