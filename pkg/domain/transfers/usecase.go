package transfers

import (
	"fmt"

	"exemplo.com/pkg/domain/entities"
)

//var transfersMap = make(map[int]entities.Transfer)
//var transferNumber = 0

type TransferUseCase struct {
	numberOfTransfers int
	storage           map[int]entities.Transfer
}

func NewTransferUseCase(numberOfTransfers int, storage map[int]entities.Transfer) TransferUseCase {
	return TransferUseCase{
		numberOfTransfers: numberOfTransfers,
		storage:           storage,
	}
}

func (tu *TransferUseCase) MakeTransfer(transfer entities.Transfer) (entities.Transfer, error) {

	transfer, err := entities.NewTransfer(transfer.Id, transfer.AccountOriginId, transfer.AccountDestinationId, transfer.Amount)

	if err != nil {
		return entities.Transfer{}, fmt.Errorf("err to create an new transfer")
	}

	transfer.Id = tu.numberOfTransfers
	tu.storage[tu.numberOfTransfers] = transfer
	tu.numberOfTransfers++

	return transfer, nil
}
