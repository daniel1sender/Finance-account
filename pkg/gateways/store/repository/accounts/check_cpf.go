package accounts

import "github.com/daniel1sender/Desafio-API/pkg/gateways/store"

func (ar accountRepository) CheckCPF(cpf string) error {
	for _, value := range ar.users {
		if value.CPF == cpf {
			return store.ErrExistingCPF
		}
	}
	return nil
}
