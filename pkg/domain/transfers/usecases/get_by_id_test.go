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

func TestTransfersUseCase_GetByID(t *testing.T) {
	transferRepository := transfers_storage.NewStorage(Db)
	accountRepository := accounts_storage.NewStorage(Db)
	ctx := context.Background()
	transferUsecase := NewUseCase(transferRepository, accountRepository)
	account1, _ := entities.NewAccount("Jonh Doe", "12345678910", "123", 10)
	account2, _ := entities.NewAccount("Jonh Doe", "12345678911", "123", 10)
	
	t.Run("should return a list of transfers", func(t *testing.T) {
		transfer, _ := entities.NewTransfer(account1.ID, account2.ID, 1)
		transferRepository.Insert(ctx, transfer)
		transfersList, err := transferUsecase.GetByID(ctx, transfer.AccountOriginID)

		assert.NotEmpty(t, transfersList)
		assert.Nil(t, err)
	})

	t.Run("should return a empty list of transfers and a error", func(t *testing.T) {
		tests.DeleteAllTransfers(Db)
		transfer, _ := entities.NewTransfer(account1.ID, account2.ID, 1)
		transfersList, err := transferUsecase.GetByID(ctx, transfer.AccountOriginID)

		assert.Empty(t, transfersList)
		assert.NotNil(t, err)
	})
}
