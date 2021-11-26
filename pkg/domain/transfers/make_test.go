package transfers

import (
	"errors"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	transfers_repository "github.com/daniel1sender/Desafio-API/pkg/gateways/store/files/transfers"
)

func TestAccountUseCase_Make(t *testing.T) {

	t.Run("should return a transfer", func(t *testing.T) {

		//storage := transfers.NewStorage()
		//transferUsecase := NewUseCase(storage)
		storage := transfers_repository.NewStorage()
		transferUsecase := NewUseCase(storage)
		amount := 10
		originID := "1"
		destinationID := "2"

		makeTransfer, err := transferUsecase.Make(originID, destinationID, amount)

		if makeTransfer == (entities.Transfer{}) {
			t.Errorf("expected a transfer but got '%+v'", makeTransfer)
		}

		if err != nil {
			t.Errorf("expected no error but got '%s'", err)
		}

	})

	t.Run("should return a blank transfer when the transfer isn't created", func(*testing.T) {

		//storage := transfers.NewStorage()
		//transferUseCase := NewUseCase(storage)
		storage := transfers_repository.NewStorage()
		transferUsecase := NewUseCase(storage)
		amount := 0
		originID := "1"
		destinationID := "2"

		makeTransfer, err := transferUsecase.Make(originID, destinationID, amount)

		if makeTransfer != (entities.Transfer{}) {
			t.Errorf("expected a blank transfer but got '%+v'", makeTransfer)
		}

		if !errors.Is(err, ErrCreatingNewTransfer) {
			t.Errorf("expected '%s' but got '%s'", ErrCreatingNewTransfer, err)
		}

	})
}
