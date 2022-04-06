package login

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type UseCaseMock struct {
	Token string
	Error error
}

func (m *UseCaseMock) Login(ctx context.Context, cpf, accountSecret string) (string, error) {
	return m.Token, m.Error
}

func (m *UseCaseMock) ValidateToken(ctx context.Context, tokenString string) (entities.Claims, error) {
	return entities.Claims{}, m.Error
}
