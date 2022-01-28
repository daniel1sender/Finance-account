package usecases

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (au AccountUseCase) GetAll(ctx context.Context) ([]entities.Account, error) {
	users, err := au.storage.GetAll(ctx)
	if err != nil {
		return []entities.Account{}, err
	}
	return users, nil
}
