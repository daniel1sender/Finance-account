package accounts

import (
	"testing"

	"exemplo.com/pkg/domain/entities"
	"exemplo.com/pkg/store"
)

func TestAccountUseCase_GetBalanceByID(t *testing.T) {

	t.Run("should return an account when id is found", func(t *testing.T) {

		storage := store.NewAccountStorage()
		AccountUseCase := NewAccountUseCase(storage)
		account, err := entities.NewAccount("John Doe", "11111111030", "123", 10)
		if err != nil {
			t.Error("Err should be nil if account was successfully created")
		}
		storage.UpdateStorage(account.ID, account)

		getBalance, err := AccountUseCase.GetBalanceByID(account.ID)

		if getBalance == 0 {
			t.Errorf("Balance account %d should be different from 0", getBalance)
		}

		if err != nil {
			t.Errorf("Err should be nil but it is %s", err)
		}

	})

	t.Run("should return a blank account when id isn't found", func(t *testing.T) {

		storage := store.NewAccountStorage()
		AccountUseCase := NewAccountUseCase(storage)
		account, err := entities.NewAccount("John Doe", "11111111030", "123", 10)
		if err != nil {
			t.Error("error should be nil if account was successfully created")
		}

		getBalance, err := AccountUseCase.GetBalanceByID(account.ID)

		if getBalance != 0 {
			t.Errorf("balance Account should be 0 but it is %d", getBalance)
		}

		if err == nil {
			t.Errorf("error should be different from nil but it is %s", err)
		}

	})

}
