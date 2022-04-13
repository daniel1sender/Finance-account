package transfers

import (
	"context"
	"errors"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

var (
	ErrEmptyList         = errors.New("empty list of transfers")
	ErrTransfersNotFound = errors.New("transfers not found from the account")
)

type UseCase interface {
	Make(ctx context.Context, originID, destinationID string, amount int) (entities.Transfer, error)
	ListByID(ctx context.Context, accountID string) ([]entities.Transfer, error)
}
