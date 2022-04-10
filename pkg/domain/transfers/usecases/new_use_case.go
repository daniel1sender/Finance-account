package usecases

import (
	"errors"

	"github.com/daniel1sender/Desafio-API/pkg/domain/transfers"
)

var (
	ErrBalanceLessZero = errors.New("balance account cannot be less than zero")
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
