package transfers

import (
	"exemplo.com/pkg/store/transfers"
)

type TransferUseCase struct {
	storage transfers.TransferStorage
}

func NewTransferUseCase(storage transfers.TransferStorage) TransferUseCase {
	return TransferUseCase{
		storage: storage,
	}
}
