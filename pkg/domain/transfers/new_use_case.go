package transfers

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/transfers"
)

type TransferUseCase struct {
	storage transfers.TransferStorage
}

func NewUseCase(storage transfers.TransferStorage) TransferUseCase {
	return TransferUseCase{
		storage: storage,
	}
}

type UseCase interface {
	Make(originID, destinationID int, amount int) (entities.Transfer, error)
}
