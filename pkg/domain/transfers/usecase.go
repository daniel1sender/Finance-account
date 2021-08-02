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
	if transfer.Amount <= 0 {
		return entities.Transfer{}, fmt.Errorf("amount equal zero")
	}
	if transfer.Account_origin_id == transfer.Account_destinantion_id {
		return entities.Transfer{}, fmt.Errorf("transfer is to the same id")
	}
	transfer.Id = tu.numberOfTransfers
	tu.storage[tu.numberOfTransfers] = transfer
	tu.numberOfTransfers++

	return transfer, nil
}
