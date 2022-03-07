package loginUseCases

import (
	"context"
)

func (l LoginUseCase) CheckToken(ctx context.Context, token string) error {
	return l.LoginStorage.CheckToken(ctx, token)
}
