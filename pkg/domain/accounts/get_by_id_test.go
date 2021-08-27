package accounts

import (
	"errors"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/store/accounts"
)

func TestAccountUseCase_GetById(t *testing.T) {

	t.Run("should return an account when the searched account is found", func(t *testing.T) {

		storage := accounts.NewStorage()
		AccountUseCase := NewUseCase(storage)

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}

		storage.Upsert(account.ID, account)
		GetAccountByID, err := AccountUseCase.GetByID(account.ID)

		if GetAccountByID == (entities.Account{}) {
			t.Errorf("expected an account but got %+v", GetAccountByID)
		}

		if err != nil {
			t.Errorf("expected error equal nil but got '%s'", err)
		}

	})

	t.Run("should return an empty account and a error message when account don't exist", func(t *testing.T) {

		storage := accounts.NewStorage()
		AccountUseCase := NewUseCase(storage)

		//passando qualquer id
		GetAccountByID, err := AccountUseCase.GetByID("account.ID")

		if GetAccountByID != (entities.Account{}) {
			t.Errorf("expected empty account but got %+v", GetAccountByID)
		}

		if !errors.Is(err, accounts.ErrIDNotFound) {
			t.Errorf("expected '%s' but got '%s'", accounts.ErrIDNotFound, err)
		}

	})

}