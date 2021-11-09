package transfers

import (
	"errors"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	accounts_storage "github.com/daniel1sender/Desafio-API/pkg/gateways/store/accounts"
	transfers_storage "github.com/daniel1sender/Desafio-API/pkg/gateways/store/transfers"
)

func TestAccountUseCase_Make(t *testing.T) {

	t.Run("should return a transfer", func(t *testing.T) {

		transferStorage := transfers_storage.NewStorage()
		accountStorage := accounts_storage.NewStorage()
		transferUsecase := NewUseCase(transferStorage, accountStorage)
		amount := 10
		originID := "1"
		destinationID := "2"
		originAccount := entities.Account{ID: originID, Balance: 20}
		destinationAccount := entities.Account{ID: destinationID, Balance: 20}
		accountStorage.Upsert(originID, originAccount)
		accountStorage.Upsert(destinationID, destinationAccount)

		makeTransfer, err := transferUsecase.Make(originID, destinationID, amount)

		if makeTransfer == (entities.Transfer{}) {
			t.Errorf("expected a transfer but got '%+v'", makeTransfer)
		}

		if makeTransfer.AccountOriginID != originAccount.ID {
			t.Errorf("expected '%s' but got '%s'", originAccount.ID, makeTransfer.AccountOriginID)
		}

		if makeTransfer.AccountDestinationID != destinationAccount.ID {
			t.Errorf("expected '%s' but got '%s'", destinationAccount.ID, makeTransfer.AccountDestinationID)
		}

		if makeTransfer.Amount != amount {
			t.Errorf("expected '%d' but got '%d'", amount, makeTransfer.Amount)
		}

		if makeTransfer.CreatedAt.IsZero() {
			t.Error("expected a time different from zero")
		}

		if err != nil {
			t.Errorf("expected no error but got '%s'", err)
		}

	})

	t.Run("should return a empty transfer when amount is less or equal zero", func(*testing.T) {

		transferStorage := transfers_storage.NewStorage()
		accountStorage := accounts_storage.NewStorage()
		transferUsecase := NewUseCase(transferStorage, accountStorage)
		amount := 0
		originID := "1"
		destinationID := "2"
		originAccount := entities.Account{ID: originID, Balance: 20}
		destinationAccount := entities.Account{ID: destinationID, Balance: 20}
		accountStorage.Upsert(originID, originAccount)
		accountStorage.Upsert(destinationID, destinationAccount)

		makeTransfer, err := transferUsecase.Make(originID, destinationID, amount)

		if makeTransfer != (entities.Transfer{}) {
			t.Errorf("expected a empty transfer but got '%+v'", makeTransfer)
		}

		if !errors.Is(err, entities.ErrAmountLessOrEqualZero) {
			t.Errorf("expected '%s' but got '%s'", entities.ErrAmountLessOrEqualZero, err)
		}

	})

	t.Run("should return an empty transfer and an error message when the origin account doesn't have sufficient funds", func(t *testing.T) {
		transferStorage := transfers_storage.NewStorage()
		accountStorage := accounts_storage.NewStorage()
		transferUseCase := NewUseCase(transferStorage, accountStorage)
		amount := 10
		originID := "1"
		destinationID := "2"

		originAccount := entities.Account{ID: originID, Balance: 0}
		accountStorage.Upsert(originID, originAccount)

		makeTransfer, err := transferUseCase.Make(originID, destinationID, amount)

		if makeTransfer != (entities.Transfer{}) {
			t.Errorf("expected a empty transfer but got '%+v'", makeTransfer)
		}

		if !errors.Is(err, ErrInsufficientFunds) {
			t.Errorf("expected '%s' but got '%s'", ErrInsufficientFunds, err)
		}

	})

	t.Run("should return a empty transfer and a error message when the account id is not found", func(t *testing.T) {

		transferStorage := transfers_storage.NewStorage()
		accountStorage := accounts_storage.NewStorage()
		transferUsecase := NewUseCase(transferStorage, accountStorage)
		amount := 0
		originID := "1"
		destinationID := "2"

		makeTransfer, err := transferUsecase.Make(originID, destinationID, amount)

		if makeTransfer != (entities.Transfer{}) {
			t.Errorf("expected a empty transfer but got '%+v'", makeTransfer)
		}

		if !errors.Is(err, accounts_storage.ErrIDNotFound) {
			t.Errorf("expected '%s' but got '%s'", accounts_storage.ErrIDNotFound, err)
		}

		originAccount := entities.Account{ID: originID, Balance: 20}
		accountStorage.Upsert(originID, originAccount)
		makeTransfer, err = transferUsecase.Make(originID, destinationID, amount)

		if makeTransfer != (entities.Transfer{}) {
			t.Errorf("expected a empty transfer but got '%+v'", makeTransfer)
		}

		if !errors.Is(err, accounts_storage.ErrIDNotFound) {
			t.Errorf("expected '%s' but got '%s'", accounts_storage.ErrIDNotFound, err)
		}

	})

}
