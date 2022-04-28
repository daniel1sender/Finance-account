package usecases

import (
	"context"
	"testing"

	accounts_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/transfers"
	"github.com/daniel1sender/Desafio-API/pkg/tests"
	"github.com/stretchr/testify/assert"
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
		account, _ := entities.NewAccount(name, cpf, secret, balance)
		accountsRespository.Upsert(ctx, account)

		err := accountUseCase.updateBalance(ctx, account.ID, 20.0)

		assert.Nil(t, err)
	})

	t.Run("should return an empty account and an error when the account doesn't exist", func(t *testing.T) {
		id := "247eade0-53b2-4ee5-9a3b-0da9afba5700"
		tests.DeleteAllAccounts(Db)

		err := accountUseCase.updateBalance(ctx, id, 20.0)

		assert.NotNil(t, err)
		assert.Equal(t, err, accounts_usecase.ErrAccountNotFound)
	})
	tests.DeleteAllAccounts(Db)
	tests.DeleteAllTransfers(Db)
}
