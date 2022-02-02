package transfers

import (
	"errors"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
)

var (
	ErrInsufficientFunds     = errors.New("insufficient balance on account")
	ErrOriginIDNotFound      = errors.New("transfer origin account id not found")
	ErrDestinationIDNotFound = errors.New("transfer destination account id not found")
)

func (tu TransferUseCase) Make(originID, destinationID string, amount int) (entities.Transfer, error) {

	originAccountBalance, err := tu.accountStorage.GetBalanceByID(originID)
	if errors.Is(err, accounts.ErrAccountNotFound) {
		return entities.Transfer{}, ErrOriginIDNotFound
	}
	if originAccountBalance < amount {
		return entities.Transfer{}, ErrInsufficientFunds
	}

	_, err = tu.accountStorage.GetByID(destinationID)
	if errors.Is(err, accounts.ErrAccountNotFound) {
		return entities.Transfer{}, ErrDestinationIDNotFound
	}

	transfer, err := entities.NewTransfer(originID, destinationID, amount)
	if err != nil {
		return entities.Transfer{}, fmt.Errorf("error creating a transfer: %w", err)
	}

	err = tu.transferStorage.UpdateByID(transfer)
	if err != nil {
		return entities.Transfer{}, err
	}

	return transfer, nil
}
