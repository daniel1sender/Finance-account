package accounts

import (
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (ar accountRepository) GetByID(id string) (entities.Account, error) {
	account, ok := ar.users[id]
	if !ok {
		return entities.Account{}, fmt.Errorf("error finding account: %w", accounts.ErrIDNotFound)
	}
	return account, nil
}
