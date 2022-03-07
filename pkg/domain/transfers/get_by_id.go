package transfers

import (
	"context"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/domain/login"
)

func (t TransferUseCase) GetByID(ctx context.Context, accountID, token, tokenSecret string) ([]entities.Transfer, error) {
	err := login.ValidateToken(token, accountID, tokenSecret)
	if err != nil {
		return []entities.Transfer{}, fmt.Errorf("error while validate token: %v", err)
	}
	transfers, err := t.transferStorage.GetByID(ctx, accountID)
	if err != nil {
		return []entities.Transfer{}, err
	}
	return transfers, nil
}
	