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

		account := entities.Account{}
		accountStorage.Upsert(originID, account)
		accountStorage.Upsert(destinationID, account)

		makeTransfer, err := transferUsecase.Make(originID, destinationID, amount)

		if makeTransfer == (entities.Transfer{}) {
			t.Errorf("expected a transfer but got '%+v'", makeTransfer)
		}

		if err != nil {
			t.Errorf("expected no error but got '%s'", err)
		}

	})

	t.Run("should return a blank transfer when the transfer isn't created", func(*testing.T) {

		transferStorage := transfers_storage.NewStorage()
		accountStorage := accounts_storage.NewStorage()
		transferUsecase := NewUseCase(transferStorage, accountStorage)
		amount := 0
		originID := "1"
		destinationID := "2"

		account := entities.Account{}
		accountStorage.Upsert(originID, account)
		accountStorage.Upsert(destinationID, account)

		makeTransfer, err := transferUsecase.Make(originID, destinationID, amount)

		if makeTransfer != (entities.Transfer{}) {
			t.Errorf("expected a blank transfer but got '%+v'", makeTransfer)
		}

		if !errors.Is(err, ErrCreatingNewTransfer) {
			t.Errorf("expected '%s' but got '%s'", ErrCreatingNewTransfer, err)
		}

	})

	t.Run("should return a blank transfer and a error message when the origin ID is not found", func(t *testing.T) {

		transferStorage := transfers_storage.NewStorage()
		accountStorage := accounts_storage.NewStorage()
		transferUsecase := NewUseCase(transferStorage, accountStorage)
		amount := 0
		originID := "1"
		destinationID := "2"

		account := entities.Account{}
		accountStorage.Upsert(destinationID, account)

		makeTransfer, err := transferUsecase.Make(originID, destinationID, amount)

		if makeTransfer != (entities.Transfer{}) {
			t.Errorf("expected a blank transfer but got '%+v'", makeTransfer)
		}

		if !errors.Is(err, accounts_storage.ErrIDNotFound) {
			t.Errorf("expected '%s' but got '%s'", accounts_storage.ErrIDNotFound, err)
		}

	})

	t.Run("should return a blank transfer and a error message when the destination ID is not found", func(t *testing.T) {

		transferStorage := transfers_storage.NewStorage()
		accountStorage := accounts_storage.NewStorage()
		transferUsecase := NewUseCase(transferStorage, accountStorage)
		amount := 0
		originID := "1"
		destinationID := "2"

		account := entities.Account{}
		accountStorage.Upsert(originID, account)

		makeTransfer, err := transferUsecase.Make(originID, destinationID, amount)

		if makeTransfer != (entities.Transfer{}) {
			t.Errorf("expected a blank transfer but got '%+v'", makeTransfer)
		}

		if !errors.Is(err, accounts_storage.ErrIDNotFound) {
			t.Errorf("expected '%s' but got '%s'", accounts_storage.ErrIDNotFound, err)
		}

	})
}
