package transfers

import (
	"errors"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

var (
	ErrCreatingNewTransfer = errors.New("error when creating a transfer")
)

func (tu TransferUseCase) Make(originID, destinationID string, amount int) (entities.Transfer, error) {

	_, err := tu.accountStorage.GetByID(originID)
	if err != nil {
		return entities.Transfer{}, fmt.Errorf("origin ID not found: %w", err)
	}
	_, err = tu.accountStorage.GetByID(destinationID)
	if err != nil {
		return entities.Transfer{}, fmt.Errorf("destination ID not found: %w", err)
	}

	transfer, err := entities.NewTransfer(originID, destinationID, amount)

	if err != nil {
		return entities.Transfer{}, ErrCreatingNewTransfer
	}

	tu.transferStorage.UpdateByID(transfer.ID, transfer)

	return transfer, nil
}
