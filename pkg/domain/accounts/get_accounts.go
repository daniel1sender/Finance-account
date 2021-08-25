package accounts

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (au AccountUseCase) GetAccounts() []entities.Account {
	return au.storage.GetAccounts()
}
