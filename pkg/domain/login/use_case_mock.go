package login

import "context"

type UseCaseMock struct {
	Token string
	Error error
}

func (m *UseCaseMock) Login(ctx context.Context, cpf, accountSecret string) (string, error) {
	return m.Token, m.Error
}
