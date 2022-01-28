package accounts

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type UseCaseMock struct {
	Balance int
	Error   error
	Account entities.Account
	List    []entities.Account
}

func (m *UseCaseMock) GetBalanceByID(ctx context.Context, id string) (int, error) {
	return m.Balance, m.Error
}

func (m *UseCaseMock) Create(ctx context.Context, name, cpf, secret string, balance int) (entities.Account, error) {
	return m.Account, m.Error
}

func (m *UseCaseMock) GetByID(ctx context.Context, id string) (entities.Account, error) {
	panic("not implemented")
}

func (m *UseCaseMock) GetAll(ctx context.Context) ([]entities.Account, error) {
	return m.List, m.Error
}

func (m *UseCaseMock) UpdateBalance(ctx context.Context, id string, balance int) error {
	panic("not implemented")
}
