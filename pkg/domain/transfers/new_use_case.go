package transfers

import (
	"context"
	"errors"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

var (
	ErrBalanceLessZero = errors.New("balance account cannot be less than zero")
)

type TransferUseCase struct {
	transferStorage Repository
	accountStorage  AccountRepository
}

func NewUseCase(transferStorage Repository, accountStorage AccountRepository) TransferUseCase {
	return TransferUseCase{
		transferStorage: transferStorage,
		accountStorage:  accountStorage,
	}
}

type UseCase interface {
	Make(ctx context.Context, originID, destinationID string, amount int) (entities.Transfer, error)
}

type Repository interface {
	Insert(ctx context.Context, transfer entities.Transfer) error
}

type AccountRepository interface {
	GetBalanceByID(ctx context.Context, id string) (int, error)
	GetByID(ctx context.Context, id string) (entities.Account, error)
	Upsert(ctx context.Context, account entities.Account) error
}
