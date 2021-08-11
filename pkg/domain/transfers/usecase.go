package transfers

import (
	"errors"

	"exemplo.com/pkg/domain/entities"
)

var (
	ErrToCallNewTransfer = errors.New("error to call function NewTransfer")
)

type TransferUseCase struct {
	storage map[string]entities.Transfer
}

func NewTransferUseCase(storage map[string]entities.Transfer) TransferUseCase {
	return TransferUseCase{
		storage: storage,
	}
}

func (tu *TransferUseCase) MakeTransfer(originID, destinationID int, amount int) (entities.Transfer, error) {

	transfer, err := entities.NewTransfer(originID, destinationID, amount)

	if err != nil {
		return entities.Transfer{}, ErrToCallNewTransfer
	}

	tu.storage[transfer.ID] = transfer

	return transfer, nil
}
