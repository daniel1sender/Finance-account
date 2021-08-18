package transfers

import (
	"exemplo.com/pkg/store"
)

type TransferUseCase struct {
	storage store.TransferStorage
}

func NewTransferUseCase(storage store.TransferStorage) TransferUseCase {
	return TransferUseCase{
		storage: storage,
	}
}
