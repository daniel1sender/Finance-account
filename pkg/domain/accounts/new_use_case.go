package accounts

import (
	"github.com/daniel1sender/Desafio-API/pkg/store/accounts"
)

type AccountUseCase struct {
	storage accounts.AccountStorage
}

func NewUseCase(storage accounts.AccountStorage) AccountUseCase {
	return AccountUseCase{
		storage: storage,
	}
}
