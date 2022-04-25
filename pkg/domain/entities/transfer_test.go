package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTransfer(t *testing.T) {

	t.Run("should successfully return a transfer", func(t *testing.T) {

		amount := 10
		originID := "1"
		destinationID := "2"
		transfer, err := NewTransfer(originID, destinationID, amount)

		assert.Nil(t, err)
		assert.Equal(t, transfer.Amount, amount)
		assert.Equal(t, transfer.AccountOriginID, originID)
		assert.Equal(t, transfer.AccountDestinationID, destinationID)
		assert.NotEmpty(t, transfer.CreatedAt)
	})

	t.Run("should return an empty transfer and an error when amount is less or equal zero", func(t *testing.T) {

		amount := 0
		originID := "1"
		destinationID := "2"
		transfer, err := NewTransfer(originID, destinationID, amount)

		assert.Equal(t, err, ErrAmountLessOrEqualZero)
		assert.Empty(t, transfer)
	})

	t.Run("should return an empty transfer and an error when transfer is to the same account", func(t *testing.T) {

		amount := 10
		originID := "1"
		destinationID := "1"
		transfer, err := NewTransfer(originID, destinationID, amount)

		assert.Equal(t, err, ErrSameAccountTransfer)
		assert.Empty(t, transfer)
	})
}
