package login

import (
	"context"
	"errors"
)

var (
	ErrTokenNotFound = errors.New("token not found")
	ErrInvalidToken  = errors.New("invalid token found")
)

type UseCase interface {
	Login(ctx context.Context, cpf, accountSecret string) (string, error)
}
