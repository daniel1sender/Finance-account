package accounts

import (
	"exemplo.com/pkg/store"
)

type AccountUseCase struct {
	storage store.AccountStorage
}

func NewAccountUseCase(storage store.AccountStorage) AccountUseCase {
	return AccountUseCase{
		storage: storage,
	}
}
