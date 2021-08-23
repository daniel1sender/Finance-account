package accounts

import (
	"exemplo.com/pkg/domain/entities"
)

func (au AccountUseCase) GetAccounts() []entities.Account {
	return au.storage.GetAccounts()
}
