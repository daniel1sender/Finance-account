package accounts

import "github.com/daniel1sender/Desafio-API/pkg/gateways/store"

func (ar AccountRepository) GetBalanceByID(id string) (int, error) {
	_, ok := ar.Users[id]
	if !ok {
		return 0, store.ErrIDNotFound
	}
	balance := ar.Users[id].Balance
	return balance, nil
}
