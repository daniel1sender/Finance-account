package transfers

import "github.com/daniel1sender/Desafio-API/pkg/domain/entities"

func (t TransferStorage) UpdateTransferStorage(id string, transfer entities.Transfer) {
	t.storage[id] = transfer
}
