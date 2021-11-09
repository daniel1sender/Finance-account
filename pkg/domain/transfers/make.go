package transfers

import (
	"errors"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

var (
	ErrCreatingNewTransfer = errors.New("error when creating a transfer")
	ErrInsufficientFunds   = errors.New("insufficient balance on account")
)

func (tu TransferUseCase) Make(originID, destinationID string, amount int) (entities.Transfer, error) {

	originAccountBalance, err := tu.accountStorage.GetBalanceByID(originID)
	if err != nil {
		return entities.Transfer{}, fmt.Errorf("error getting balance account by id: %w", err)
	}
	if originAccountBalance < amount {
		return entities.Transfer{}, ErrInsufficientFunds
	}

	_, err = tu.accountStorage.GetByID(destinationID)
	if err != nil {
		return entities.Transfer{}, fmt.Errorf("error finding an account: %w", err)
	}

	transfer, err := entities.NewTransfer(originID, destinationID, amount)
	if err != nil {
		return entities.Transfer{}, fmt.Errorf("error when creating a transfer: %w", err)
	}

	tu.transferStorage.UpdateByID(transfer.ID, transfer)
	tu.accountStorage.UpdateBalance(originID, destinationID, amount)

	return transfer, nil
}
