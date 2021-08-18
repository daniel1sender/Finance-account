package transfers

import (
	"errors"

	"exemplo.com/pkg/domain/entities"
)

var (
	ErrToCallNewTransfer = errors.New("error to call function NewTransfer")
)

func (tu *TransferUseCase) MakeTransfer(originID, destinationID int, amount int) (entities.Transfer, error) {

	transfer, err := entities.NewTransfer(originID, destinationID, amount)

	if err != nil {
		return entities.Transfer{}, ErrToCallNewTransfer
	}

	tu.storage.UpdateTransferStorage(transfer.ID, transfer)

	return transfer, nil
}
