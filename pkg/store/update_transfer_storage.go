package store

import "exemplo.com/pkg/domain/entities"

func (t TransferStorage) UpdateTransferStorage(id string, transfer entities.Transfer) {
	t.storage[id] = transfer
}
