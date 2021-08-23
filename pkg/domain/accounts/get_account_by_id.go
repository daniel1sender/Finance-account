package accounts

import (
	"exemplo.com/pkg/domain/entities"
)

func (au AccountUseCase) GetAccountByID(id string) (entities.Account, error) {
	return au.storage.GetAccountByID(id)
}
