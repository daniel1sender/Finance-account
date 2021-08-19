package accounts

import (
	"exemplo.com/pkg/store/accounts"
)

type AccountUseCase struct {
	storage accounts.AccountStorage
}

func NewAccountUseCase(storage accounts.AccountStorage) AccountUseCase {
	return AccountUseCase{
		storage: storage,
	}
}
