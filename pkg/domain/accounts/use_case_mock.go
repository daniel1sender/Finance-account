package accounts

import (
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type UseCaseMock struct {
	Balance int
	Error   error
	Account entities.Account
	List    []entities.Account
}

func (m *UseCaseMock) GetBalanceByID(id string) (int, error) {
	return m.Balance, m.Error
}

func (m *UseCaseMock) Create(name, cpf, secret string, balance int) (entities.Account, error) {

	account, err := entities.NewAccount(name, cpf, secret, balance)
	if err != nil {
		return entities.Account{}, fmt.Errorf("%s: %w", ErrCreatingNewAccount, err)
	}

	m.List = append(m.List, account)

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
