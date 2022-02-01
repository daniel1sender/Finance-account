package transfers

import (
	"context"
	"errors"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"

	accounts_storage "github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/accounts"
	transfers_storage "github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/transfers"

)

func TestAccountUseCase_Make(t *testing.T) {
	transferRepository := transfers_storage.NewStorage(Db)
	accountRepository := accounts_storage.NewStorage(Db)
	ctx := context.Background()

	t.Run("should return a transfer", func(t *testing.T) {

		transferUsecase := NewUseCase(transferRepository, accountRepository)
		amount := 10

		name := "John Doe"
		cpf1 := "11111111030"
		cpf2 := "11111111031"
		secret := "123"
		balance := 10

		originAccount, err := entities.NewAccount(name, cpf1, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}
		destinationAccount, err := entities.NewAccount(name, cpf2, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}
		accountRepository.Upsert(ctx, originAccount)
		accountRepository.Upsert(ctx, destinationAccount)

		makeTransfer, err := transferUsecase.Make(ctx, originAccount.ID, destinationAccount.ID, amount)

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
		DeleteAll(Db)
	})

	t.Run("should return an empty transfer and a error when amount is less or equal zero", func(*testing.T) {

		transferUsecase := NewUseCase(transferRepository, accountRepository)
		amount := 0

		name := "John Doe"
		cpf1 := "11111111031"
		cpf2 := "11111111032"
		secret := "123"
		balance := 10

		originAccount, err := entities.NewAccount(name, cpf1, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}
		destinationAccount, err := entities.NewAccount(name, cpf2, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}
		accountRepository.Upsert(ctx, originAccount)
		accountRepository.Upsert(ctx, destinationAccount)

		makeTransfer, err := transferUsecase.Make(ctx, originAccount.ID, destinationAccount.ID, amount)

		if makeTransfer != (entities.Transfer{}) {
			t.Errorf("expected an empty transfer but got '%+v'", makeTransfer)
		}

		if !errors.Is(err, entities.ErrAmountLessOrEqualZero) {
			t.Errorf("expected '%s' but got '%s'", entities.ErrAmountLessOrEqualZero, err)
		}
		DeleteAll(Db)
	})

	t.Run("should return an empty transfer and an error when the origin account doesn't have sufficient funds", func(t *testing.T) {

		transferUseCase := NewUseCase(transferRepository, accountRepository)
		amount := 10
		name := "John Doe"
		cpf1 := "11111111031"
		cpf2 := "11111111032"
		secret := "123"
		balance := 0


		originAccount, err := entities.NewAccount(name, cpf1, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}
		destinationAccount, err := entities.NewAccount(name, cpf2, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}

		accountRepository.Upsert(ctx, originAccount)

		makeTransfer, err := transferUseCase.Make(ctx, originAccount.ID, destinationAccount.ID, amount)

		if makeTransfer != (entities.Transfer{}) {
			t.Errorf("expected an empty transfer but got '%+v'", makeTransfer)
		}

		if !errors.Is(err, ErrInsufficientFunds) {
			t.Errorf("expected '%s' but got '%s'", ErrInsufficientFunds, err)
		}
		DeleteAll(Db)
	})

	t.Run("should return an empty transfer and an error when the transfer origin account id is not found", func(t *testing.T) {

		transferUseCase := NewUseCase(transferRepository, accountRepository)
		amount := 10
		name := "John Doe"
		cpf1 := "11111111031"
		cpf2 := "11111111032"
		secret := "123"
		balance := 10

		originAccount, err := entities.NewAccount(name, cpf1, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}
		destinationAccount, err := entities.NewAccount(name, cpf2, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}

		makeTransfer, err := transferUseCase.Make(ctx, originAccount.ID, destinationAccount.ID, amount)

		if makeTransfer != (entities.Transfer{}) {
			t.Errorf("expected an empty transfer but got '%+v'", makeTransfer)
		}

		if !errors.Is(err, ErrOriginAccountNotFound) {
			t.Errorf("expected '%s' but got '%s'", ErrOriginAccountNotFound, err)
		}
	})

	t.Run("should return an empty transfer and an error when the transfer destination account id is not found", func(t *testing.T) {

		transferUsecase := NewUseCase(transferRepository, accountRepository)
		amount := 10
		name := "John Doe"
		cpf1 := "11111111031"
		cpf2 := "11111111032"
		secret := "123"
		balance := 10

		originAccount, err := entities.NewAccount(name, cpf1, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}
		destinationAccount, err := entities.NewAccount(name, cpf2, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}
		accountRepository.Upsert(ctx, originAccount)

		makeTransfer, err := transferUsecase.Make(ctx, originAccount.ID, destinationAccount.ID, amount)
		
		if makeTransfer != (entities.Transfer{}) {
			t.Errorf("expected a empty transfer but got '%+v'", makeTransfer)
		}

		if !errors.Is(err, ErrDestinationAccountNotFound) {
			t.Errorf("expected '%s' but got '%s'", ErrDestinationAccountNotFound, err)
		}

	})
	DeleteAll(Db)
}
