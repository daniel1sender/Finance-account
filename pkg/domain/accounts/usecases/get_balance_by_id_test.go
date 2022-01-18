package usecases

import (
	"context"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	accounts_repository "github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/accounts"
)

func TestAccountUseCase_GetBalanceByID(t *testing.T) {

	t.Run("should return an account balance when id is found", func(t *testing.T) {

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

		getBalance, err := accountUsecase.GetBalanceByID(account.ID)

		if getBalance == 0 {
			t.Error("expected balance account different from 0")
		}

		if err != nil {
			t.Errorf("expected no error but got '%s'", err)
		}

	})

	t.Run("should return a null account balance and an error when id isn't found", func(t *testing.T) {

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
		repository.Exec(context.Background(), "DELETE FROM accounts")

		getBalance, err := accountUsecase.GetBalanceByID(account.ID)

		if getBalance != 0 {
			t.Error("expected account balance equal zero")
		}

		if err == nil {
			t.Error("expected no error")
		}

	})

}
