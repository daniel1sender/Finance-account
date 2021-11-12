package accounts

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type AccountStorage struct {
	storage map[string]entities.Account
}

func NewStorage() AccountStorage {
	sto := make(map[string]entities.Account)
	return AccountStorage{
		storage: sto,
	}
}
