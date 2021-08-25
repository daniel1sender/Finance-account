package accounts

import (
	"errors"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

var (
	ErrExistingCPF = errors.New("cpf informed is invalid")
	ErrIDNotFound  = errors.New("account id isn't found")
)

type AccountStorage struct {
	storage map[string]entities.Account
}

func NewAccountStorage() AccountStorage {
	sto := make(map[string]entities.Account)
	return AccountStorage{
		storage: sto,
	}
}
