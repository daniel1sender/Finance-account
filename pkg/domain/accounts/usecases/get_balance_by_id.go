package usecases

import "fmt"

func (au AccountUseCase) GetBalanceByID(id string) (int, error) {
	balance, err := au.storage.GetBalanceByID(id)
	if err != nil {
		return balance, fmt.Errorf("error while getting balance account: %w", err)
	}
	return balance, err
}
