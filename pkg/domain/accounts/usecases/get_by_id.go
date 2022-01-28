package usecases

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (au AccountUseCase) GetByID(ctx context.Context, id string) (entities.Account, error) {
	account, err := au.storage.GetByID(ctx, id)
	if err != nil {
		return entities.Account{}, err
	}
	return account, err
}
