package transfers

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/transfers"
)

type TransferUseCase struct {
	transferStorage transfers.TransferStorage
	accountStorage  accounts.AccountStorage
}

func NewUseCase(transferStorage transfers.TransferStorage, accountStorage accounts.AccountStorage) TransferUseCase {
	return TransferUseCase{
		transferStorage: transferStorage,
		accountStorage:  accountStorage,
	}
}

type UseCase interface {
	Make(originID, destinationID string, amount int) (entities.Transfer, error)
}
