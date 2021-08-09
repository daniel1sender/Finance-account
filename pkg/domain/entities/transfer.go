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
	Id                   string
	AccountOriginId      int
	AccountDestinationId int
	Amount               int
	CreatedAt            time.Time
}

func NewTransfer(originId, destinationId int, amount int) (Transfer, error) {

	if amount <= 0 {
		return Transfer{}, ErrAmountLessThanZero

	}
	if originId == destinationId {
		return Transfer{}, ErrSameAccountTransfer
	}

	id := uuid.NewString()

	return Transfer{
		Id:                   id,
		AccountOriginId:      originId,
		AccountDestinationId: destinationId,
		Amount:               amount,
		CreatedAt:            time.Now().UTC(),
	}, nil

}
