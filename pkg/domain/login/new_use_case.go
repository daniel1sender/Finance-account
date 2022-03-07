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
	Auth(ctx context.Context, cpf, secret string) (string, string, error)
	CheckToken(ctx context.Context, token string) error
	GetTokenByID(ctx context.Context, id string) (string, error)
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
	CheckToken(ctx context.Context, token string) error
	Insert(ctx context.Context, token, tokenSecret string) error
	GetTokenByID(ctx context.Context, id string) (string, error)
}
