package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrAmountLessThanZero  = errors.New("amount is less then zero")
	ErrSameAccountTransfer = errors.New("transfer attempt to the same account")
)

type Transfer struct {
	ID                   string
	AccountOriginID      int
	AccountDestinationID int
	Amount               int
	CreatedAt            time.Time
}

func NewTransfer(originID, destinationID int, amount int) (Transfer, error) {

	if amount <= 0 {
		return Transfer{}, ErrAmountLessThanZero

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
