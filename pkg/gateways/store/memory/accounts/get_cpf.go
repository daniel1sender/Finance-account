package accounts

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
)

func (s AccountStorage) CheckCPF(cpf string) error {
	for _, storedAccount := range s.storage {
		if storedAccount.CPF == cpf {
			return accounts.ErrExistingCPF
		}
	}
	return nil
}
