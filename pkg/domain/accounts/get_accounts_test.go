package accounts

import (
	"testing"

	"exemplo.com/pkg/domain/entities"
	"exemplo.com/pkg/store/accounts"
)

func TestAccountUseCase_GetAccounts(t *testing.T) {

	t.Run("should return a full list of accounts", func(t *testing.T) {

		storage := accounts.NewAccountStorage()
		AccountUseCase := NewAccountUseCase(storage)
		account, err := entities.NewAccount("John Doe", "11111111030", "123", 10)
		if err != nil {
			t.Error("error should be nil if account was successfully created")
		}
		storage.UpdateStorage(account.ID, account)

		getAccounts := AccountUseCase.GetAccounts()

		if len(getAccounts) == 0 {
			t.Error("expected a full list of accounts")
		}

	})

	t.Run("should return an empty list", func(t *testing.T) {

		storage := accounts.NewAccountStorage()
		AccountUseCase := NewAccountUseCase(storage)

		getAccounts := AccountUseCase.GetAccounts()

		if len(getAccounts) != 0 {
			t.Error("expected an empty list")
		}

	})

}
