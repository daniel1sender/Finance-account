package usecases

import (
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
)

func (au AccountUseCase) GetBalanceByID(id string) (int, error) {
	balance, err := au.storage.GetBalanceByID(id)
	if err != nil {
		return balance, fmt.Errorf("%v : %w", accounts.ErrIDNotFound, err)
	}
	return balance, err
}
