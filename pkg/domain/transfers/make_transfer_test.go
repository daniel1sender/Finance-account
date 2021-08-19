package transfers

import (
	"testing"

	"exemplo.com/pkg/domain/entities"
	"exemplo.com/pkg/store/transfers"
)

func TestAccountUseCase_MakeTransfer(t *testing.T) {

	t.Run("Should return a transfer", func(t *testing.T) {

		storage := transfers.NewTransferStorage()
		TransferUsecase := NewTransferUseCase(storage)

		MakeTransfer, err := TransferUsecase.MakeTransfer(1, 2, 10)

		if MakeTransfer == (entities.Transfer{}) {
			t.Errorf("Expected a transfer but got %+v", MakeTransfer)
		}

		if err != nil {
			t.Errorf("Expected nil error but got %s", err)
		}

	})

}
