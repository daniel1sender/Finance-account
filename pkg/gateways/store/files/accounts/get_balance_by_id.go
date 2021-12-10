package accounts

import (
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
)

func (ar accountRepository) GetBalanceByID(id string) (int, error) {
	_, ok := ar.users[id]
	if !ok {
		return 0, fmt.Errorf("error finding account: %w", accounts.ErrIDNotFound)
	}
	balance := ar.users[id].Balance
	return balance, nil
}
