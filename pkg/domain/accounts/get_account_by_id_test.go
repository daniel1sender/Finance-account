package accounts

import (
	"errors"
	"testing"

	"exemplo.com/pkg/domain/entities"
	"exemplo.com/pkg/store/accounts"
)

func TestAccountUseCase_GetAccountById(t *testing.T) {

	t.Run("Should return an account when the searched account is found", func(t *testing.T) {

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
		GetAccountByID, err := AccountUseCase.GetAccountByID(account.ID)

		if GetAccountByID == (entities.Account{}) {
			t.Errorf("Expected account but got %+v", GetAccountByID)
		}

		if err != nil {
			t.Error("Expected error equal nil")
		}

	})

	t.Run("Should return an empty account and a error message when account don't exist", func(t *testing.T) {

		storage := accounts.NewAccountStorage()
		AccountUseCase := NewAccountUseCase(storage)

		//passando qualquer id
		GetAccountByID, err := AccountUseCase.GetAccountByID("account.ID")

		if GetAccountByID != (entities.Account{}) {
			t.Errorf("Expected empty account but got %+v", GetAccountByID)
		}

		if !errors.Is(err, accounts.ErrIDNotFound) {
			t.Errorf("Expected %s but got %s", accounts.ErrIDNotFound, err)
		}

	})

}
