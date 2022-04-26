package accounts

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type AccountRepository struct {
	storage map[string]entities.Account
}

func NewRepository() AccountRepository {
	sto := make(map[string]entities.Account)
	return AccountRepository{
		storage: sto,
	}
}
