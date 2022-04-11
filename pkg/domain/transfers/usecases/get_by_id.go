package usecases

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (tu TransferUseCase) GetByID(ctx context.Context, accountID string) ([]entities.Transfer, error) {
	return tu.transferStorage.GetByID(ctx, accountID)
}
