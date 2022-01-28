package usecases

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrBalanceLessZero = errors.New("balance account cannot be less than zero")
)

func (au AccountUseCase) UpdateBalance(ctx context.Context, id string, balance int) error {
	account, err := au.storage.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if balance < 0 {
		return fmt.Errorf("error updating balance account: %w", ErrBalanceLessZero)
	}

	account.Balance = balance
	err = au.storage.Upsert(ctx, account)
	if err != nil {
		return err
	}

	return nil
}
