package usecases

import (
	"errors"
	"testing"

	accounts_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	accounts_repository "github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/accounts"
)

func TestAccountUseCase_GetById(t *testing.T) {
	repository := accounts_repository.NewStorage(Db)
	accountUseCase := NewUseCase(repository)

	t.Run("should return an account when the searched account is found", func(t *testing.T) {
		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}

		repository.Upsert(account)
		getAccountByID, err := accountUseCase.GetByID(account.ID)

		if getAccountByID == (entities.Account{}) {
			t.Errorf("expected an account but got %+v", getAccountByID)
		}

		if err != nil {
			t.Errorf("expected error equal nil but got '%s'", err)
		}

	})

	t.Run("should return an empty account and a error when account don't exist", func(t *testing.T) {

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}
		DeleteAll(Db)

		getAccountByID, err := accountUseCase.GetByID(account.ID)

		if getAccountByID != (entities.Account{}) {
			t.Errorf("expected empty account but got %+v", getAccountByID)
		}

		if !errors.Is(err, accounts_usecase.ErrAccountNotFound) {
			t.Errorf("expected '%s' but got '%s'", accounts_usecase.ErrAccountNotFound, err)
		}

	})

}
