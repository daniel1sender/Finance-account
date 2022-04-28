package usecases

import (
	"context"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	accounts_repository "github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/tests"
	"github.com/stretchr/testify/assert"
)

func TestAccountUseCase_GetAll(t *testing.T) {
	repository := accounts_repository.NewStorage(Db)
	accountUsecase := NewUseCase(repository)
	ctx := context.Background()

	t.Run("should return a list of accounts", func(t *testing.T) {

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10
		account, _ := entities.NewAccount(name, cpf, secret, balance)
		repository.Upsert(ctx, account)

		getAccounts, err := accountUsecase.GetAll(ctx)
		assert.Nil(t, err)
		assert.NotEmpty(t, getAccounts)
		assert.Equal(t, getAccounts[0].Name, account.Name)
		assert.Equal(t, getAccounts[0].CPF, account.CPF)
		assert.Equal(t, getAccounts[0].Balance, account.Balance)
		assert.Equal(t, getAccounts[0].Secret, account.Secret) 
	})

	t.Run("should return an empty list", func(t *testing.T) {
		tests.DeleteAllAccounts(Db)

		getAccounts, err := accountUsecase.GetAll(ctx)
		assert.NotEqual(t, err, accounts.ErrEmptyList)
		assert.Empty(t, getAccounts)
	})
}
