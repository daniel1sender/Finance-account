package login

import "context"

func (l LoginUseCase) GetTokenByID(ctx context.Context, id string) (string, error) {
	return l.LoginStorage.GetTokenByID(ctx, id)
}
