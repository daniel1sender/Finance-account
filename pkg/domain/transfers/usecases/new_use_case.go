package usecases

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/transfers"
)

type TransferUseCase struct {
	transferStorage transfers.Repository
	accountStorage  transfers.AccountRepository
}

func NewUseCase(transferStorage transfers.Repository, accountStorage transfers.AccountRepository) TransferUseCase {
	return TransferUseCase{
		transferStorage: transferStorage,
		accountStorage:  accountStorage,
	}
}
