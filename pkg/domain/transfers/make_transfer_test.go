package transfers

import (
	"errors"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/store/transfers"
)

func TestAccountUseCase_MakeTransfer(t *testing.T) {

	t.Run("should return a transfer", func(t *testing.T) {

		storage := transfers.NewTransferStorage()
		TransferUsecase := NewTransferUseCase(storage)
		amount := 10
		originID := 1
		destinationID := 2

		MakeTransfer, err := TransferUsecase.MakeTransfer(originID, destinationID, amount)

		if MakeTransfer == (entities.Transfer{}) {
			t.Errorf("expected a transfer but got '%+v'", MakeTransfer)
		}

		if err != nil {
			t.Errorf("expected no error but got '%s'", err)
		}

	})

	t.Run("should return a blank transfer when the transfer isn't created", func(*testing.T) {

		storage := transfers.NewTransferStorage()
		transferUseCase := NewTransferUseCase(storage)
		amount := 0
		originID := 1
		destinationID := 2

		MakeTransfer, err := transferUseCase.MakeTransfer(originID, destinationID, amount)

		if MakeTransfer != (entities.Transfer{}) {
			t.Errorf("expected a blank transfer but got '%+v'", MakeTransfer)
		}

		if !errors.Is(err, ErrToCreateNewTransfer) {
			t.Errorf("expected '%s' but got '%s'", ErrToCreateNewTransfer, err)
		}

	})
}
