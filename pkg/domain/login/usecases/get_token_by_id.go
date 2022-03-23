package usecases

import "context"

func (l LoginUseCase) GetTokenByID(ctx context.Context, tokenID string) (string, error){
	return l.LoginRepository.GetTokenByID(ctx, tokenID)
}
