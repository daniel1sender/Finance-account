package transfers

import (
	"errors"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

var (
	ErrToCreateNewTransfer = errors.New("error to create a new transfer")
)

func (tu *TransferUseCase) MakeTransfer(originID, destinationID int, amount int) (entities.Transfer, error) {

	transfer, err := entities.NewTransfer(originID, destinationID, amount)

	if err != nil {
		return entities.Transfer{}, ErrToCreateNewTransfer
	}

	tu.storage.UpdateTransferStorage(transfer.ID, transfer)

	return transfer, nil
}
