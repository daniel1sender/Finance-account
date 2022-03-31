package login

import (
	"context"
	"errors"
)

var (
	ErrTokenNotFound = errors.New("token not found")
	ErrInvalidToken  = errors.New("invalid token found")
	ErrEmptySecret   = errors.New("secret informed is blanc")
	ErrInvalidCPF    = errors.New("cpf informed is invalid")
	ErrInvalidSecret = errors.New("secret informed is incorrect")
)

type UseCase interface {
	Login(ctx context.Context, cpf, accountSecret string) (string, error)
}
