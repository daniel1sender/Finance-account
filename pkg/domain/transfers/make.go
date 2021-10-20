package transfers

import (
	"errors"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

var (
	ErrCreatingNewTransfer = errors.New("error when creating a transfer")
	ErrNoOriginID          = errors.New("origin ID not found")
	ErrNoDestinationID     = errors.New("destination ID not found")
)

func (tu TransferUseCase) Make(originID, destinationID string, amount int) (entities.Transfer, error) {

	_, err := tu.accountStorage.GetByID(originID)
	if err != nil {
		return entities.Transfer{}, fmt.Errorf("origin ID %s not found", originID)
	}
	_, err = tu.accountStorage.GetByID(destinationID)
	if err != nil {
		return entities.Transfer{}, fmt.Errorf("destination ID %s not found", destinationID)
	}

	transfer, err := entities.NewTransfer(originID, destinationID, amount)

	if err != nil {
		return entities.Transfer{}, ErrCreatingNewTransfer
	}

	tu.transferStorage.UpdateByID(transfer.ID, transfer)

	return transfer, nil
}
