package accounts

import "github.com/daniel1sender/Desafio-API/pkg/gateways/store"

func (ar AccountRepository) CheckCPF(cpf string) error {

	for _, value := range ar.Users{
		if value.CPF == cpf{
			return store.ErrExistingCPF
		}
	}
	return nil
}
