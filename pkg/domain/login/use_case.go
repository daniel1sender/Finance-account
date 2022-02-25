package login

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type LoginUseCase struct {
	LoginStorage   Repository
	AccountStorage AccountRepository
	tokenSecret    string
}

type UseCase interface {
	Auth(ctx context.Context, cpf, secret string) (string, error)
}

type AccountRepository interface {
	GetByCPF(ctx context.Context, cpf string) (entities.Account, error)
}

func NewUseCase(loginStorage Repository, accountStorage AccountRepository, tokenSecret string) LoginUseCase {
	return LoginUseCase{
		LoginStorage:   loginStorage,
		AccountStorage: accountStorage,
		tokenSecret:    tokenSecret,
	}
}

type Repository interface {
	GetTokenByID(ctx context.Context, tokenID string) (string, error)
	Insert(ctx context.Context, token, tokenSecret string) error
}