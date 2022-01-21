package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrAmountLessOrEqualZero = errors.New("amount is less or equal zero")
	ErrSameAccountTransfer   = errors.New("transfer attempt to the same account")
	ErrEmptyOriginID         = errors.New("invalid origin id")
	ErrEmptyDestinationID    = errors.New("invalid destination id")
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

	if originID == "" {
		return Transfer{}, ErrEmptyOriginID
	}

	if destinationID == "" {
		return Transfer{}, ErrEmptyDestinationID
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
