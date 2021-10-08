package transfers

import (
	"errors"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/transfers"
)

func TestAccountUseCase_Make(t *testing.T) {

	t.Run("should return a transfer", func(t *testing.T) {

		storage := transfers.NewStorage()
		transferUsecase := NewUseCase(storage)
		amount := 10
		originID := 1
		destinationID := 2

		makeTransfer, err := transferUsecase.Make(originID, destinationID, amount)

		if makeTransfer == (entities.Transfer{}) {
			t.Errorf("expected a transfer but got '%+v'", makeTransfer)
		}

		if err != nil {
			t.Errorf("expected no error but got '%s'", err)
		}

	})

	t.Run("should return a empty transfer when transfer it is not created", func(*testing.T) {

		storage := transfers.NewStorage()
		transferUseCase := NewUseCase(storage)
		amount := 0
		originID := 1
		destinationID := 2

		makeTransfer, err := transferUseCase.Make(originID, destinationID, amount)

		if makeTransfer != (entities.Transfer{}) {
			t.Errorf("expected a blank transfer but got '%+v'", makeTransfer)
		}

		if !errors.Is(err, entities.ErrNegativeAmount) {
			t.Errorf("expected '%s' but got '%s'", entities.ErrNegativeAmount, err)
		}

	})
}
