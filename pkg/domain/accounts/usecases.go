package accounts

import (
	"errors"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

var (
	ErrExistingCPF = errors.New("cpf informed alredy exists")
	ErrIDNotFound  = errors.New("account id isn't found")
)

type UseCase interface {
	GetBalanceByID(id string) (int, error)
	Create(name, cpf, secret string, balance int) (entities.Account, error)
	GetByID(id string) (entities.Account, error)
	UpdateBalance(id string, balance int) error
	GetAll() []entities.Account
}

type Repository interface {
	GetAll() []entities.Account
	GetBalanceByID(id string) (int, error)
	GetByID(id string) (entities.Account, error)
	CheckCPF(cpf string) error
	Upsert(account entities.Account)
}
