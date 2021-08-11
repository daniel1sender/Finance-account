package entities

import (
	"errors"
	"testing"
)

func TestNewTransfer(t *testing.T) {

	t.Run("Should successfully return a transfer", func(t *testing.T) {

		amount := 10
		originId := 1
		destinationId := 2
		transfer, err := NewTransfer(originId, destinationId, amount)
		if err != nil {
			t.Errorf("Expected nil error but got %s", err)
		}

		if transfer.Amount != amount {
			t.Errorf("Expected amount %d but got %d", transfer.Amount, amount)
		}

		if transfer.AccountOriginId != originId {
			t.Errorf("Expected originId %d but got %d", originId, transfer.AccountOriginId)
		}

		if transfer.AccountDestinationId != destinationId {
			t.Errorf("Expected originId %d but got %d", destinationId, transfer.AccountDestinationId)
		}

		if transfer.AccountOriginId == transfer.AccountDestinationId {
			t.Error("Expected a transfer to different accounts")
		}

		if transfer.CreatedAt.IsZero() == true {
			t.Error("Expected a time different from zero")
		}

	})

	t.Run("Should return a empty transfer when amount is less or equal zero", func(t *testing.T) {

		amount := 0
		originId := 1
		destinationId := 2
		transfer, err := NewTransfer(originId, destinationId, amount)

		if !errors.Is(err, ErrAmountLessThanZero) {
			t.Errorf("Expected error %s but got %s", ErrAmountLessThanZero, err)
		}

		if transfer != (Transfer{}) {
			t.Errorf("Expected %+v but got %+v", Transfer{}, transfer)
		}

	})

	t.Run("Should return a empty transfer when transfer is to the same account", func(t *testing.T) {

		amount := 10
		originId := 1
		destinationId := 1
		transfer, err := NewTransfer(originId, destinationId, amount)

		if !errors.Is(err, ErrSameAccountTransfer) {
			t.Errorf("Expected error %s but got %s", ErrSameAccountTransfer, err)
		}

		if transfer != (Transfer{}) {
			t.Errorf("Expected empty trasnfer but got %+v", transfer)
		}

	})

}
