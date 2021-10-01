package accounts

import (
	"time"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type UseCaseMock struct {
	Balance int
	Error   error
	List    []entities.Account
}

func (m *UseCaseMock) GetBalanceByID(id string) (int, error) {
	return m.Balance, m.Error
}

func (m *UseCaseMock) Create(name, cpf, secret string, balance int) (entities.Account, error) {

	if name == "" {
		return entities.Account{}, entities.ErrInvalidName
	}

	if len(cpf) != 11 {
		return entities.Account{}, entities.ErrInvalidCPF
	}

	if secret == "" {
		return entities.Account{}, entities.ErrBlankSecret
	}

	if balance < 0 {
		return entities.Account{}, entities.ErrBalanceLessZero
	}

	account := entities.Account{Name: name, CPF: cpf, Secret: secret, Balance: balance, CreatedAt: time.Now()}

	return account, nil
}

func (m *UseCaseMock) GetByID(id string) (entities.Account, error) {
	panic("not implemented")
}

func (m *UseCaseMock) Get() []entities.Account {
	return m.List
}

func (m *UseCaseMock) UpdateBalance(id string, balance int) error {
	panic("not implemented")
}
