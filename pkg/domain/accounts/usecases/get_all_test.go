package usecases

import (
	"context"
	"testing"

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

		getAccounts := accountUsecase.GetAll()

		if len(getAccounts) == 0 {
			t.Error("expected a full list of accounts")
		}

	})

	t.Run("should return an empty list", func(t *testing.T) {
		repository := accounts_repository.NewStorage(Db)
		accountUsecase := NewUseCase(repository)
		repository.Exec(context.Background(), "DELETE FROM accounts")

		getAccounts := accountUsecase.GetAll()

		if len(getAccounts) != 0 {
			t.Error("expected an empty list")
		}

	})

}
