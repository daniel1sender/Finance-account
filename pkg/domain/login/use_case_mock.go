package login

import (
	"context"
)

type LoginUseCaseMock struct {
	Token       string
	AccountID   string
	TokenSecret string
}

func (l *LoginUseCaseMock) CheckToken(ctx context.Context, token string) error {
	return nil
}

func (l *LoginUseCaseMock) Auth(ctx context.Context, cpf, secret string) (string, string, error) {
	panic("not implemented")
}

func (l *LoginUseCaseMock) GetTokenByID(ctx context.Context, id string) (string, error) {
	return "", nil
}
