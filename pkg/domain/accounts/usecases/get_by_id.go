package usecases

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (au AccountUseCase) GetByID(id string) (entities.Account, error) {
	account, err := au.storage.GetByID(id)
	if err != nil {
		return entities.Account{}, err
	}
	return account, err
}
