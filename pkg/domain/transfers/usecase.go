package transfers

import (
	"fmt"

	"exemplo.com/pkg/domain/entities"
)

type TransferUseCase struct {
	storage map[string]entities.Transfer
}

func NewTransferUseCase(storage map[string]entities.Transfer) TransferUseCase {
	return TransferUseCase{
		storage: storage,
	}
}

func (tu *TransferUseCase) MakeTransfer(originId, destinationId int, amount float64) (entities.Transfer, error) {

	transfer, err := entities.NewTransfer(originId, destinationId, amount)

	if err != nil {
		return entities.Transfer{}, fmt.Errorf("err to create a new transfer")
	}

	tu.storage[transfer.Id] = transfer

	return transfer, nil
}
