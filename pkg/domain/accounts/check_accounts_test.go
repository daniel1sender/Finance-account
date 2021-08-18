package accounts

import (
	"testing"

	"exemplo.com/pkg/domain/entities"
	"exemplo.com/pkg/store"
)

func TestAccountUseCase_CheckAccounts(t *testing.T) {

	t.Run("should return nil when accounts have already been created", func(t *testing.T) {

		storage := store.NewAccountStorage()
		AccountUseCase := NewAccountUseCase(storage)
		account, err := entities.NewAccount("John Doe", "11111111030", "123", 10)
		if err != nil {
			t.Error("error should be nil if account was successfully created")
		}

		storage.UpdateStorage(account.ID, account)

		CheckAccountsError := AccountUseCase.CheckAccounts(account.ID)

		if CheckAccountsError != nil {
			t.Error("expected nil when account exists")
		}

	})

	t.Run("should return an error message when id isn't found", func(t *testing.T) {

		storage := store.NewAccountStorage()
		AccountUseCase := NewAccountUseCase(storage)
		account, err := entities.NewAccount("John Doe", "11111111030", "123", 10)
		if err != nil {
			t.Error("error should be nil if account was successfully created")
		}

		CheckAccountsError := AccountUseCase.CheckAccounts(account.ID)

		if CheckAccountsError == nil {
			t.Error("expected a error message")
		}

	})

}
