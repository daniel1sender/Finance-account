package login

import (
	"context"
)

type LoginUseCase struct {
	AccountStorage AccountRepository
	tokenSecret    string
}

type UseCase interface {
	Auth(ctx context.Context, cpf, secret string) (string, error)
}

func NewUseCase(accountStorage AccountRepository, tokenSecret string) LoginUseCase {
	return LoginUseCase{
		AccountStorage: accountStorage,
		tokenSecret:    tokenSecret,
	}
}
