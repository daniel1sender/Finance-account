package transfers

import "github.com/daniel1sender/Desafio-API/pkg/domain/entities"

func (t TransferRepository) UpdateByID(transfer entities.Transfer) error {
	t.storage[transfer.ID] = transfer
	return nil
}
