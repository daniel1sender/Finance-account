package usecases

import (
	"errors"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	accounts_repository "github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/accounts"
)

func TestAccountUseCase_GetAll(t *testing.T) {

	t.Run("should return a full list of accounts", func(t *testing.T) {
		repository := accounts_repository.NewStorage(Db)
		accountUsecase := NewUseCase(repository)
		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}

		repository.Upsert(account)

		getAccounts, err := accountUsecase.GetAll()

		if err != nil {
			t.Errorf("expected null error but got %v", err)
		}

		if len(getAccounts) == 0 {
			t.Error("expected a full list of accounts")
		}

	})

	t.Run("should return an empty list", func(t *testing.T) {
		repository := accounts_repository.NewStorage(Db)
		accountUsecase := NewUseCase(repository)
		repository.DeleteAll()

		getAccounts, err := accountUsecase.GetAll()

		if !errors.Is(err, accounts.ErrEmptyList) {
			t.Errorf("expected %v but got %v", accounts.ErrEmptyList, err)
		}

		if len(getAccounts) != 0 {
			t.Error("expected an empty list")
		}

	})

}
