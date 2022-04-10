package transfers

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type UseCase interface {
	Make(ctx context.Context, originID, destinationID string, amount int) (entities.Transfer, error)
}
