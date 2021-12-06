package transfers

import (
	"errors"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

var (
	ErrCreatingNewTransfer = errors.New("error when creating a transfer")
)

func (tu TransferUseCase) Make(originID, destinationID string, amount int) (entities.Transfer, error) {

	transfer, err := entities.NewTransfer(originID, destinationID, amount)

	if err != nil {
		return entities.Transfer{}, ErrCreatingNewTransfer
	}

	err = tu.storage.UpdateByID(transfer)
	if err != nil{
		return entities.Transfer{}, err
	}

	return transfer, nil
}
