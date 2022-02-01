package transfers

import (
	"context"
)

func (au TransferUseCase) UpdateBalance(ctx context.Context, id string, amount int) error {
	account, err := au.accountStorage.GetByID(ctx, id)
	if err != nil {
		return err
	}

	balance := account.Balance + amount
	account.Balance = balance
	err = au.accountStorage.Upsert(ctx, account)
	if err != nil {
		return err
	}

	return nil
}
