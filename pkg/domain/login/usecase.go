package login

import (
	"context"
	"errors"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

var (
	ErrTokenNotFound      = errors.New("token not found")
	ErrInvalidToken       = errors.New("invalid token found")
	ErrInvalidSecret      = errors.New("secret informed is incorrect")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type UseCase interface {
	Login(ctx context.Context, cpf, accountSecret string) (string, error)
	ValidateToken(ctx context.Context, tokenString string) (entities.Claims, error)
}
