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

func TestAccountUseCase_GetById(t *testing.T) {
	repository := accounts_repository.NewStorage(Db)
	accountUseCase := NewUseCase(repository)
	ctx := context.Background()

	t.Run("should return an account when the account id it is found", func(t *testing.T) {
		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10
		account, _ := entities.NewAccount(name, cpf, secret, balance)
		repository.Upsert(ctx, account)

		getAccountByID, err := accountUseCase.GetByID(ctx, account.ID)

		assert.Nil(t, err)
		assert.NotEmpty(t, getAccountByID)
	})

	t.Run("should return an empty account and an error when the account doesn't exist", func(t *testing.T) {

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, _ := entities.NewAccount(name, cpf, secret, balance)
		tests.DeleteAllAccounts(Db)

		getAccountByID, err := accountUseCase.GetByID(ctx, account.ID)

		assert.Empty(t, getAccountByID)
		assert.Equal(t, err, accounts_usecase.ErrAccountNotFound)
	})

}
