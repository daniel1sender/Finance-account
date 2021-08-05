package transfers

import (
	"testing"

	"exemplo.com/pkg/domain/entities"
)

func TestAccountUseCase_MakeTransfer(t *testing.T) {

	t.Run("Should return a transfer", func(t *testing.T) {

		storage := make(map[string]entities.Transfer)
		TransferUsecase := NewTransferUseCase(storage)

		MakeTransfer, err := TransferUsecase.MakeTransfer(1, 2, 10)

		if MakeTransfer == (entities.Transfer{}) {
			t.Errorf("Expected a transfer but got %+v", MakeTransfer)
		}

		if err != nil {
			t.Error("Expecte null error")
		}

	})

	t.Run("should return an empty transfer when origin id is equal destination id", func(t *testing.T) {

		storage := make(map[string]entities.Transfer)
		TransferUseCase := NewTransferUseCase(storage)

		MakeTransfer, err := TransferUseCase.MakeTransfer(1, 1, 0)

		compareTransfer(t, MakeTransfer, (entities.Transfer{}))

		if err == nil {
			t.Error("Expected err message")
		}

	})

	t.Run("Should return an empty transfer when amount is less or equal zero", func(t *testing.T) {

		storage := make(map[string]entities.Transfer)
		TransferUseCase := NewTransferUseCase(storage)

		MakeTransfer, err := TransferUseCase.MakeTransfer(1, 2, -10)

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
