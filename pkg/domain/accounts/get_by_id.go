package accounts

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (au AccountUseCase) GetByID(id string) (entities.Account, error) {
	return au.storage.GetByID(id)
}
