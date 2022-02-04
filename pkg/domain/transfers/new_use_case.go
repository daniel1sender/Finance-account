package transfers

import (
	"context"
	"errors"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
)

var (
	ErrBalanceLessZero = errors.New("balance account cannot be less than zero")
)

type TransferUseCase struct {
	transferStorage Repository
	accountStorage  accounts.Repository
}

func NewUseCase(transferStorage Repository, accountStorage accounts.Repository) TransferUseCase {
	return TransferUseCase{
		transferStorage: transferStorage,
		accountStorage:  accountStorage,
	}
}

type UseCase interface {
	Make(ctx context.Context, originID, destinationID string, amount int) (entities.Transfer, error)
	UpdateBalance(ctx context.Context, id string, balance int) error
}

type Repository interface {
	Insert(ctx context.Context, transfer entities.Transfer) error
}
