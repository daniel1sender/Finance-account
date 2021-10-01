package accounts

import (
	"time"

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

	account := entities.Account{Name: m.Account.Name, CPF: m.Account.CPF, Secret: m.Account.Secret, Balance: m.Account.Balance, CreatedAt: time.Now()}

	return account, m.Error
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
