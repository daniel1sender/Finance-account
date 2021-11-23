package transfers

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type TransferUseCase struct {
	storage Repository
}

func NewUseCase(storage Repository) TransferUseCase {
	return TransferUseCase{
		storage: storage,
	}
}

type UseCase interface {
	Make(originID, destinationID string, amount int) (entities.Transfer, error)
}

type Repository interface {
	UpdateByID(transfer entities.Transfer)error
}
