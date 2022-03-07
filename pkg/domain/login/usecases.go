package login

import (
	"context"
	"errors"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

var (
	ErrTokenNotFound = errors.New("token not found")
)

type UseCase interface {
	Auth(ctx context.Context, cpf, secret string) (string, string, error)
	CheckToken(ctx context.Context, token string) error
	GetTokenByID(ctx context.Context, id string) (string, error)
}

type AccountRepository interface {
	GetByCPF(ctx context.Context, cpf string) (entities.Account, error)
}

type Repository interface {
	CheckToken(ctx context.Context, token string) error
	Insert(ctx context.Context, token, tokenSecret string) error
	GetTokenByID(ctx context.Context, id string) (string, error)
}
