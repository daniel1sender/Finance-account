package usecases

import (
	"errors"
	"testing"

	accounts_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	accounts_repository "github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/accounts"
)

func TestAccountUseCase_UpdateBalance(t *testing.T) {
	repository := accounts_repository.NewStorage(Db)
	accountUseCase := NewUseCase(repository)

	t.Run("should return an account and null error when account was updated", func(t *testing.T) {
		name := "John Doe"
		cpf := "12345678010"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}

		repository.Upsert(account)

		updateAccountError := accountUseCase.UpdateBalance(account.ID, 20.0)

		if updateAccountError != nil {
			t.Errorf("expected no error but got '%s'", updateAccountError)
		}

	})
	
	t.Run("should return an empty account an error when account don't exists", func(t *testing.T) {
		name := "John Doe"
		cpf := "11111111031"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}
		DeleteAll(Db)

		err = accountUseCase.UpdateBalance(account.ID, 20.0)

		if errors.Is(err, accounts_usecase.ErrAccountNotFound) {
			t.Errorf("expected '%s' but got '%s'", accounts_usecase.ErrAccountNotFound, err)
		}

	})

	t.Run("should return an empty account and an error when balance account is less than zero", func(t *testing.T) {
		name := "John Doe"
		cpf := "11111111039"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}

		repository.Upsert(account)

		err = accountUseCase.UpdateBalance(account.ID, -10)

		if !errors.Is(err, ErrBalanceLessZero) {
			t.Errorf("expected '%s' but got '%s'", ErrBalanceLessZero, err)
		}

	})
	DeleteAll(Db)
}
