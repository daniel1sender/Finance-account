package accounts

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (au AccountUseCase) Get() []entities.Account {
	return au.storage.Get()
}
