package transfers

import (
	"context"
)

func (au TransferUseCase) updateBalance(ctx context.Context, id string, amount int) error {
	account, err := au.accountStorage.GetByID(ctx, id)
	if err != nil {
		return err
	}

	account.Balance += amount
	err = au.accountStorage.Upsert(ctx, account)
	if err != nil {
		return err
	}

	return nil
}
