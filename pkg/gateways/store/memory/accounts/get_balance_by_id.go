package accounts

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
)

func (s AccountRepository) GetBalanceByID(id string) (int, error) {
	for key, value := range s.storage {
		if value.ID == id {
			balance := s.storage[key].Balance
			return balance, nil
		}
	}
	return 0, accounts.ErrAccountNotFound
}
