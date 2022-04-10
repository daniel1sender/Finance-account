package usecases

import (
	"context"
	"errors"
	"testing"

	accounts_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/transfers"
	"github.com/daniel1sender/Desafio-API/pkg/tests"
)

func TestTranferUseCase_UpdateBalance(t *testing.T) {
	transfersRespository := transfers.NewStorage(Db)
	accountsRespository := accounts.NewStorage(Db)
	accountUseCase := NewUseCase(transfersRespository, accountsRespository)
	ctx := context.Background()

	t.Run("should return an account and null error when account was updated", func(t *testing.T) {
		name := "John Doe"
		cpf := "12345678010"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}

		accountsRespository.Upsert(ctx, account)

		updateAccountError := accountUseCase.updateBalance(ctx, account.ID, 20.0)

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
		tests.DeleteAllAccounts(Db)

		err = accountUseCase.updateBalance(ctx, account.ID, 20.0)

		if !errors.Is(err, accounts_usecase.ErrAccountNotFound) {
			t.Errorf("expected '%s' but got '%s'", accounts_usecase.ErrAccountNotFound, err)
		}

	})
	tests.DeleteAllAccounts(Db)
	tests.DeleteAllTransfers(Db)
}
