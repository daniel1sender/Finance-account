package usecases

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (au AccountUseCase) GetAll() ([]entities.Account, error) {
	users, err := au.storage.GetAll()
	if err != nil {
		return []entities.Account{}, err
	}
	return users, nil
}
