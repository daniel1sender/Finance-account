package usecases

import (
	"context"
	"testing"

	accounts_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/accounts"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	accounts_repository "github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/tests"
	"github.com/stretchr/testify/assert"
)

func TestAccountUseCase_Create(t *testing.T) {
	repository := accounts_repository.NewStorage(Db)
	accountUsecase := NewUseCase(repository)
	ctx := context.Background()

	t.Run("should successfully create an account and return it", func(t *testing.T) {

		name := "John Doe"
		cpf := "11111111032"
		secret := "123"
		balance := 10

		account := entities.Account{Name: name, CPF: cpf, Secret: secret, Balance: balance}

		createdAccount, err := accountUsecase.Create(ctx, account.Name, account.CPF, account.Secret, account.Balance)

		assert.Nil(t, err)
		assert.NotEmpty(t, createdAccount)
		assert.Equal(t, createdAccount.Name, account.Name)
		assert.Equal(t, createdAccount.CPF, account.CPF)
		assert.Equal(t, createdAccount.Balance, account.Balance)
	})

	t.Run("should return an empty account and an error when trying to create account with already created cpf account", func(t *testing.T) {

		name := "John Doe"
		cpf := "12111111038"
		secret := "123"
		balance := 10

		_, _ = accountUsecase.Create(ctx, name, cpf, secret, balance)

		createdAccount, err := accountUsecase.Create(ctx, name, cpf, secret, balance)

		assert.Equal(t, err, accounts_usecase.ErrExistingCPF)
		assert.Empty(t, createdAccount)
	})
	tests.DeleteAllAccounts(Db)
}
