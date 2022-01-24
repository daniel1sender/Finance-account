package usecases

import (
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (au AccountUseCase) GetByID(id string) (entities.Account, error) {
	account, err := au.storage.GetByID(id)
	if err != nil {
		return entities.Account{}, fmt.Errorf("%w: %v", accounts.ErrIDNotFound, err)
	}
	return account, err
}
