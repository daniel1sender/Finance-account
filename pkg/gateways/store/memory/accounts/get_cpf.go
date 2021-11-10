package accounts

import "github.com/daniel1sender/Desafio-API/pkg/gateways/store"

func (s AccountStorage) CheckCPF(cpf string) error {
	for _, storedAccount := range s.storage {
		if storedAccount.CPF == cpf {
			return store.ErrExistingCPF
		}
	}
	return nil
}
