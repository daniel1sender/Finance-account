package transfers

import (
	"context"
	"errors"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

var (
	ErrInsufficientFunds          = errors.New("insufficient balance on account")
	ErrOriginAccountNotFound      = errors.New("transfer origin account not found")
	ErrDestinationAccountNotFound = errors.New("transfer destination account not found")
)

func (tu TransferUseCase) Make(ctx context.Context, originID, destinationID string, amount int) (entities.Transfer, error) {

	originAccountBalance, err := tu.accountStorage.GetBalanceByID(ctx, originID)
	if err != nil {
		if errors.Is(err, accounts.ErrAccountNotFound) {
			return entities.Transfer{}, fmt.Errorf("%w: %s", ErrOriginAccountNotFound, accounts.ErrAccountNotFound.Error())
		} else {
			return entities.Transfer{}, fmt.Errorf("error to get balance account: %s", err.Error())
		}
	}

	if originAccountBalance < amount {
		return entities.Transfer{}, ErrInsufficientFunds
	}

	_, err = tu.accountStorage.GetByID(ctx, destinationID)
	if err != nil {
		if errors.Is(err, accounts.ErrAccountNotFound) {
			return entities.Transfer{}, fmt.Errorf("%w: %s", ErrDestinationAccountNotFound, accounts.ErrAccountNotFound.Error())
		} else {
			return entities.Transfer{}, fmt.Errorf("error to get balance account: %s", err.Error())
		}
	}

	transfer, err := entities.NewTransfer(originID, destinationID, amount)
	if err != nil {
		return entities.Transfer{}, fmt.Errorf("error while creating a transfer: %w", err)
	}

	err = tu.transferStorage.UpdateByID(ctx, transfer)
	if err != nil {
		return entities.Transfer{}, fmt.Errorf("error while inserting the transfer: %w", err)
	}

	err = tu.UpdateBalance(ctx, originID, -amount)
	if err != nil {
		return entities.Transfer{}, fmt.Errorf("error while updating the balance of origin account: %w", err)
	}
	err = tu.UpdateBalance(ctx, destinationID, amount)
	if err != nil {
		return entities.Transfer{}, fmt.Errorf("error while updating the balance of destination account: %w", err)
	}

	return transfer, nil
}
