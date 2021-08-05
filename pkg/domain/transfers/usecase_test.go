package transfers

import (
	"testing"
	"time"

	"exemplo.com/pkg/domain/entities"
)

func TestAccountUseCase_MakeTransfer(t *testing.T) {

	t.Run("Should return a transfer", func(t *testing.T) {

		transfer := entities.Transfer{
			Id:                   0,
			AccountOriginId:      1,
			AccountDestinationId: 2,
			Amount:               10,
			CreatedAt:            time.Time{},
		}
		storage := make(map[int]entities.Transfer)
		TransferUsecase := NewTransferUseCase(0, storage)
		storage[0] = transfer

		MakeTransfer, err := TransferUsecase.MakeTransfer(transfer)

		if MakeTransfer == (entities.Transfer{}) {
			t.Errorf("Expected a transfer but got %+v", MakeTransfer)
		}

		if err != nil {
			t.Error("Expecte null error")
		}

	})

	t.Run("should return an empty transfer when origin id is equal destination id", func(t *testing.T) {

		transfer := entities.Transfer{Id: 0, AccountOriginId: 0, AccountDestinationId: 0, Amount: 20}
		storage := make(map[int]entities.Transfer)
		TransferUseCase := NewTransferUseCase(0, storage)

		MakeTransfer, err := TransferUseCase.MakeTransfer(transfer)

		compareTransfer(t, MakeTransfer, (entities.Transfer{}))

		if err == nil {
			t.Error("Expected err message")
		}

	})

	t.Run("Should return an empty transfer when amount is less or equal zero", func(t *testing.T) {

		transfer := entities.Transfer{Id: 0, AccountOriginId: 0, AccountDestinationId: 0, Amount: 20}
		storage := make(map[int]entities.Transfer)
		TransferUseCase := NewTransferUseCase(0, storage)

		MakeTransfer, err := TransferUseCase.MakeTransfer(transfer)

		compareTransfer(t, MakeTransfer, (entities.Transfer{}))

		if err == nil {
			t.Error("Expected error message")
		}

	})

}

func compareTransfer(t *testing.T, resultado, esperado entities.Transfer) {
	t.Helper()

	if resultado != esperado {
		t.Error("Expected an empty transfer")
	}
}
