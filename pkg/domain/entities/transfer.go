package entities

import (
	"fmt"
	"time"
)

type Transfer struct {
	Id                   int
	AccountOriginId      int
	AccountDestinationId int
	Amount               float64
	CreatedAt            time.Time
}

func NewTransfer(id, originId, destinationId int, amount float64) (Transfer, error) {

	if amount <= 0 {
		return Transfer{}, fmt.Errorf("amount equal zero")
	}

	if originId == destinationId {
		return Transfer{}, fmt.Errorf("transfer is to the same id")
	}

	return Transfer{
		Id:                   id,
		AccountOriginId:      originId,
		AccountDestinationId: destinationId,
		Amount:               amount,
		CreatedAt:            time.Time{},
	}, nil

}
