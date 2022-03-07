package loginUseCases

import "context"

func (l LoginUseCase) GetTokenByID(ctx context.Context, AccountID string) (string, error) {
	return l.LoginStorage.GetTokenByID(ctx, AccountID)
}
