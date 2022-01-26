package usecases

import (
	"errors"
	"fmt"
)

var (
	ErrBalanceLessZero = errors.New("balance account cannot be less than zero")
)

func (au AccountUseCase) UpdateBalance(id string, balance int) error {
	account, err := au.storage.GetByID(id)
	if err != nil {
		return err
	}
	
	if balance < 0 {
		return fmt.Errorf("error updating balance account: %w", ErrBalanceLessZero)
	}

	account.Balance = balance
	err = au.storage.Upsert(account)
	if err != nil {
		return  err
	}

	return nil
}
