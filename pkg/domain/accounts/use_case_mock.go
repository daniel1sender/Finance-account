package accounts

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type UseCaseMock struct {
	Balance int
	Error error
}

func (m *UseCaseMock) GetBalanceByID(id string) (int, error) {
	return m.Balance, m.Error
}

func (m *UseCaseMock) Create(name, cpf, secret string, balance int) (entities.Account, error) {
	panic("not implemented")
}

func (m *UseCaseMock) GetByID(id string) (entities.Account, error) {
	panic("not implemented")
}

func (m *UseCaseMock) Get() []entities.Account {
	panic("not implemented")
}

func (m *UseCaseMock) UpdateBalance(id string, balance int) error {
	panic("not implemented")
}
