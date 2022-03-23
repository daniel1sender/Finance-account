package usecases

import "context"

func (l LoginUseCase) CheckToken(ctx context.Context, token string) error{
	return l.LoginRepository.CheckToken(ctx, token)
}
