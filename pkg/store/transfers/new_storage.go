package transfers

import "exemplo.com/pkg/domain/entities"

type TransferStorage struct {
	storage map[string]entities.Transfer
}

func NewTransferStorage() TransferStorage {
	sto := make(map[string]entities.Transfer)
	return TransferStorage{
		storage: sto,
	}
}
