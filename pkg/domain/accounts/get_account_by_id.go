package accounts

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (au AccountUseCase) GetAccountByID(id string) (entities.Account, error) {
	return au.storage.GetAccountByID(id)
}
