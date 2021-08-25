package accounts

import (
	"errors"
)

var (
	ErrBalanceLessOrEqualZero = errors.New("balance account less or equal zero")
)

func (au AccountUseCase) UpdateAccountBalance(id string, balance int) error {
	account, err := au.GetAccountByID(id)
	if err != nil {
		return err
	}
	if balance < 0 {
		return ErrBalanceLessOrEqualZero
	}

	account.Balance = balance
	au.storage.UpdateByID(account.ID, account)

	return nil
}
