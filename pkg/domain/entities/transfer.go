package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrAmountLessOrEqualZero = errors.New("amount is less or equal zero")
	ErrSameAccountTransfer   = errors.New("transfer attempt to the same account")
	ErrInsufficientFunds   = errors.New("insufficient balance on account")
)

type Transfer struct {
	ID                   string
	AccountOriginID      string
	AccountDestinationID string
	Amount               int
	CreatedAt            time.Time
}

func NewTransfer(originID, destinationID string, amount int) (Transfer, error) {

	if amount <= 0 {
		return Transfer{}, ErrAmountLessOrEqualZero

	}
	if originID == destinationID {
		return Transfer{}, ErrSameAccountTransfer
	}

	id := uuid.NewString()

	return Transfer{
		ID:                   id,
		AccountOriginID:      originID,
		AccountDestinationID: destinationID,
		Amount:               amount,
		CreatedAt:            time.Now().UTC(),
	}, nil

}
