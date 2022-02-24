package login

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type LoginUseCase struct {
	AccountStorage AccountRepository
	tokenSecret    string
}

type UseCase interface {
	Auth(ctx context.Context, cpf, secret string) (string, error)
}

type AccountRepository interface {
	GetByCPF(ctx context.Context, cpf string) (entities.Account, error)
}

func NewUseCase(accountStorage AccountRepository, tokenSecret string) LoginUseCase {
	return LoginUseCase{
		AccountStorage: accountStorage,
		tokenSecret:    tokenSecret,
	}
}
