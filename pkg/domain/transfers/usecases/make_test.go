package usecases

import (
	"context"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/tests"
	"github.com/stretchr/testify/assert"

	accounts_storage "github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/accounts"
	transfers_storage "github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/transfers"
)

func TestTransfersUseCase_Create(t *testing.T) {
	transferRepository := transfers_storage.NewStorage(Db)
	accountRepository := accounts_storage.NewStorage(Db)
	ctx := context.Background()

	t.Run("should return a transfer", func(t *testing.T) {

		transferUsecase := NewUseCase(transferRepository, accountRepository)
		amount := 10

		name := "John Doe"
		cpf1 := "11111111030"
		cpf2 := "11111111031"
		secret := "123"
		balance := 10
		originAccount, _ := entities.NewAccount(name, cpf1, secret, balance)
		destinationAccount, _ := entities.NewAccount(name, cpf2, secret, balance)
		accountRepository.Upsert(ctx, originAccount)
		accountRepository.Upsert(ctx, destinationAccount)

		transfer, err := transferUsecase.Make(ctx, originAccount.ID, destinationAccount.ID, amount)

		assert.NotEmpty(t, transfer)
		assert.Equal(t, transfer.AccountOriginID, originAccount.ID)
		assert.Equal(t, transfer.AccountDestinationID, destinationAccount.ID)
		assert.Equal(t, transfer.Amount, amount)
		assert.NotEmpty(t, transfer.CreatedAt)
		assert.Nil(t, err)
		tests.DeleteAllAccounts(Db)
	})

	t.Run("should return an empty transfer and an error when amount is less or equal zero", func(*testing.T) {

		transferUsecase := NewUseCase(transferRepository, accountRepository)
		amount := 0

		name := "John Doe"
		cpf1 := "11111111031"
		cpf2 := "11111111032"
		secret := "123"
		balance := 10
		originAccount, _ := entities.NewAccount(name, cpf1, secret, balance)
		destinationAccount, _ := entities.NewAccount(name, cpf2, secret, balance)
		accountRepository.Upsert(ctx, originAccount)
		accountRepository.Upsert(ctx, destinationAccount)

		transfer, err := transferUsecase.Make(ctx, originAccount.ID, destinationAccount.ID, amount)

		assert.Empty(t, transfer)
		assert.NotEqual(t, err, entities.ErrAmountLessOrEqualZero)
		tests.DeleteAllAccounts(Db)
	})

	t.Run("should return an empty transfer and an error when the origin account doesn't have sufficient funds", func(t *testing.T) {

		transferUseCase := NewUseCase(transferRepository, accountRepository)
		amount := 10
		name := "John Doe"
		cpf1 := "11111111031"
		cpf2 := "11111111032"
		secret := "123"
		balance := 0
		originAccount, _ := entities.NewAccount(name, cpf1, secret, balance)
		destinationAccount, _ := entities.NewAccount(name, cpf2, secret, balance)
		accountRepository.Upsert(ctx, originAccount)

		transfer, err := transferUseCase.Make(ctx, originAccount.ID, destinationAccount.ID, amount)

		assert.Empty(t, transfer)
		assert.Equal(t, err, ErrInsufficientFunds)
		tests.DeleteAllAccounts(Db)
	})

	t.Run("should return an empty transfer and an error when the transfer origin account id is not found", func(t *testing.T) {

		transferUseCase := NewUseCase(transferRepository, accountRepository)
		amount := 10
		name := "John Doe"
		cpf1 := "11111111031"
		cpf2 := "11111111032"
		secret := "123"
		balance := 10
		originAccount, _ := entities.NewAccount(name, cpf1, secret, balance)
		destinationAccount, _ := entities.NewAccount(name, cpf2, secret, balance)

		transfer, err := transferUseCase.Make(ctx, originAccount.ID, destinationAccount.ID, amount)

		assert.Empty(t, transfer)
		assert.NotEqual(t, err, ErrOriginAccountNotFound)
	})

	t.Run("should return an empty transfer and an error when the transfer destination account id is not found", func(t *testing.T) {

		transferUsecase := NewUseCase(transferRepository, accountRepository)
		amount := 10
		name := "John Doe"
		cpf1 := "11111111031"
		cpf2 := "11111111032"
		secret := "123"
		balance := 10
		originAccount, _ := entities.NewAccount(name, cpf1, secret, balance)
		destinationAccount, _ := entities.NewAccount(name, cpf2, secret, balance)
		accountRepository.Upsert(ctx, originAccount)

		transfer, err := transferUsecase.Make(ctx, originAccount.ID, destinationAccount.ID, amount)

		assert.Empty(t, transfer)
		assert.NotEqual(t, err, ErrDestinationAccountNotFound)
	})
	tests.DeleteAllAccounts(Db)
	tests.DeleteAllTransfers(Db)
}
