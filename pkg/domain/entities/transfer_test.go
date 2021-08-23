package entities

import (
	"errors"
	"testing"
)

func TestNewTransfer(t *testing.T) {

	t.Run("Should successfully return a transfer", func(t *testing.T) {

		amount := 10
		originID := 1
		destinationID := 2
		transfer, err := NewTransfer(originID, destinationID, amount)
		if err != nil {
			t.Errorf("Expected nil error but got '%s'", err)
		}

		if transfer.Amount != amount {
			t.Errorf("Expected amount '%d' but got '%d'", transfer.Amount, amount)
		}

		if transfer.AccountOriginID != originID {
			t.Errorf("Expected originId '%d' but got '%d'", originID, transfer.AccountOriginID)
		}

		if transfer.AccountDestinationID != destinationID {
			t.Errorf("Expected originId '%d' but got '%d'", destinationID, transfer.AccountDestinationID)
		}

		if transfer.CreatedAt.IsZero() == true {
			t.Error("Expected a time different from zero")
		}

	})

	t.Run("Should return a empty transfer when amount is less or equal zero", func(t *testing.T) {

		amount := 0
		originID := 1
		destinationID := 2
		transfer, err := NewTransfer(originID, destinationID, amount)

		if !errors.Is(err, ErrAmountLessOrEqualZero) {
			t.Errorf("Expected error '%s' but got '%s'", ErrAmountLessOrEqualZero, err)
		}

		if transfer != (Transfer{}) {
			t.Errorf("Expected '%+v' but got '%+v'", Transfer{}, transfer)
		}

	})

	t.Run("Should return a empty transfer when transfer is to the same account", func(t *testing.T) {

		amount := 10
		originID := 1
		destinationID := 1
		transfer, err := NewTransfer(originID, destinationID, amount)

		if !errors.Is(err, ErrSameAccountTransfer) {
			t.Errorf("Expected error '%s' but got '%s'", ErrSameAccountTransfer, err)
		}

		if transfer != (Transfer{}) {
			t.Errorf("Expected blank transfer but got '%+v'", transfer)
		}

	})

}
