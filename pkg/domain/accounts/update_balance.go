package accounts

import (
	"errors"
)

var (
	ErrBalanceLessZero = errors.New("balance account cannot be less than zero")
)

func (au AccountUseCase) UpdateBalance(id string, balance int) error {
	account, err := au.GetByID(id)
	if err != nil {
		return err
	}
	if balance < 0 {
		return ErrBalanceLessZero
	}

	account.Balance = balance
	au.storage.Upsert(account.ID, account)

	return nil
}
