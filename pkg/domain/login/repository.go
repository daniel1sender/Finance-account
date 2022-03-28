package login

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type AccountRepository interface {
	GetByCPF(ctx context.Context, cpf string) (entities.Account, error)
}

type Repository interface {
	GetTokenByID(ctx context.Context, tokenID string) (string, error)
	Insert(ctx context.Context, claims entities.Claims, token string) error
}
