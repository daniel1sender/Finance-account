package transfers

import (
	"context"
	"errors"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

var (
	ErrEmptyList = errors.New("got empty list of transfers")
)

type UseCase interface {
	Make(ctx context.Context, originID, destinationID string, amount int) (entities.Transfer, error)
	ListByAccountID(ctx context.Context, accountID string) ([]entities.Transfer, error)
}
