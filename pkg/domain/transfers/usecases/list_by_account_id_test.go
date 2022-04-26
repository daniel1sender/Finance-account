package usecases

import (
	"context"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	accounts_storage "github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/accounts"
	transfers_storage "github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/transfers"
	"github.com/daniel1sender/Desafio-API/pkg/tests"
	"github.com/stretchr/testify/assert"
)

func TestTransfersUseCase_ListByAccountID(t *testing.T) {
	transferRepository := transfers_storage.NewRepository(Db)
	accountRepository := accounts_storage.NewRepository(Db)
	ctx := context.Background()
	transferUsecase := NewUseCase(transferRepository, accountRepository)
	account1, _ := entities.NewAccount("Jonh Doe", "12345678910", "123", 10)
	account2, _ := entities.NewAccount("Jonh Doe", "12345678911", "123", 10)

	t.Run("should return a list of transfers", func(t *testing.T) {
		transfer, _ := entities.NewTransfer(account1.ID, account2.ID, 1)
		transferRepository.Insert(ctx, transfer)
		transfersList, err := transferUsecase.ListByAccountID(ctx, transfer.AccountOriginID)

		assert.Equal(t, transfersList[0].ID, transfer.ID)
		assert.Equal(t, transfersList[0].AccountOriginID, transfer.AccountOriginID)
		assert.Equal(t, transfersList[0].AccountDestinationID, transfer.AccountDestinationID)
		assert.Equal(t, transfersList[0].Amount, transfer.Amount)
		assert.NotEmpty(t, transfersList[0].CreatedAt)
		assert.NotEmpty(t, transfersList)
		assert.Nil(t, err)
	})

	t.Run("should return an empty list of transfers and an error", func(t *testing.T) {
		tests.DeleteAllTransfers(Db)
		transfer, _ := entities.NewTransfer(account1.ID, account2.ID, 1)
		transfersList, err := transferUsecase.ListByAccountID(ctx, transfer.AccountOriginID)

		assert.Empty(t, transfersList)
		assert.NotNil(t, err)
	})
}
