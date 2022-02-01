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
	if errors.Is(err, accounts.ErrAccountNotFound) {
		return entities.Transfer{}, ErrOriginAccountNotFound
	}
	if originAccountBalance < amount {
		return entities.Transfer{}, ErrInsufficientFunds
	}

	_, err = tu.accountStorage.GetByID(ctx, destinationID)
	if errors.Is(err, accounts.ErrAccountNotFound) {
		return entities.Transfer{}, ErrDestinationAccountNotFound
	}

	transfer, err := entities.NewTransfer(originID, destinationID, amount)
	if err != nil {
		return entities.Transfer{}, fmt.Errorf("error while creating a transfer: %w", err)
	}

	err = tu.transferStorage.UpdateByID(ctx, transfer)
	if err != nil {
		return entities.Transfer{}, err
	}

	err = tu.UpdateBalance(ctx, originID, -amount)
	if err != nil {
		return entities.Transfer{}, err
	}
	err = tu.UpdateBalance(ctx, destinationID, amount)

	if err != nil {
		return entities.Transfer{}, err
	}

	return transfer, nil
}
