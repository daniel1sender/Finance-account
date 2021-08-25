package transfers

import "github.com/daniel1sender/Desafio-API/pkg/domain/entities"

type TransferStorage struct {
	storage map[string]entities.Transfer
}

func NewTransferStorage() TransferStorage {
	sto := make(map[string]entities.Transfer)
	return TransferStorage{
		storage: sto,
	}
}
