package accounts

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/accounts"
)

type AccountUseCase struct {
	storage accounts.AccountStorage
}

func NewUseCase(storage accounts.AccountStorage) AccountUseCase {
	return AccountUseCase{
		storage: storage,
	}
}

type UseCase interface {
	GetBalanceByID(id string) (int, error)
	Create(name, cpf, secret string, balance int) (entities.Account, error)
 	GetByID(id string) (entities.Account, error)
	UpdateBalance(id string, balance int) error 
	GetAll() []entities.Account
}