package accounts

import (
	"errors"
)

var (
	ErrBalanceLessZero = errors.New("balance account cannot be less than zero")
)

func (au AccountUseCase) UpdateAccountBalance(id string, balance int) error {
	account, err := au.GetAccountByID(id)
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
