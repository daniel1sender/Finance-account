package login

import (
	"context"
	"errors"
)

var (
	ErrTokenNotFound = errors.New("token not found")
	ErrInvalidToken  = errors.New("invalid token found")
	ErrEmptySecret   = errors.New("empty secret was informed")
	ErrInvalidCPF    = errors.New("cpf informed is invalid")
	ErrInvalidSecret = errors.New("secret informed is incorrect")
	ErrInvalidCredentials= errors.New("invalid credentials")
)

type UseCase interface {
	Login(ctx context.Context, cpf, accountSecret string) (string, error)
}
