package transfers

import "github.com/daniel1sender/Desafio-API/pkg/domain/entities"

func (t TransferStorage) UpdateByID(transfer entities.Transfer) {
	t.storage[transfer.ID] = transfer
}
