package usecases

import "context"

func (au AccountUseCase) GetBalanceByID(ctx context.Context, id string) (int, error) {
	balance, err := au.storage.GetBalanceByID(ctx, id)
	if err != nil {
		return 0, err
	}
	return balance, nil
}
