package accounts

import "github.com/daniel1sender/Desafio-API/pkg/gateways/store"

func (ar accountRepository) GetBalanceByID(id string) (int, error) {
	_, ok := ar.users[id]
	if !ok {
		return 0, store.ErrIDNotFound
	}
	balance := ar.users[id].Balance
	return balance, nil
}
