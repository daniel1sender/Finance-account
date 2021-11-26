package usecases

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (au AccountUseCase) GetAll() []entities.Account {
	return au.storage.GetAll()
}
