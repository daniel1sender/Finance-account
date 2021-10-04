package transfers

import (
	"errors"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

//checar depois esse erro deve ser uma sentinela
var (
	ErrCreatingNewTransfer = errors.New("error when creating a transfer")
)

func (tu *TransferUseCase) Make(originID, destinationID int, amount int) (entities.Transfer, error) {

	transfer, err := entities.NewTransfer(originID, destinationID, amount)

	if err != nil {
		return entities.Transfer{}, fmt.Errorf("error while creating transfer: %w", err)
	}

	tu.storage.UpdateByID(transfer.ID, transfer)

	return transfer, nil
}
