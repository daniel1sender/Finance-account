package accounts

import (
	"testing"

	"exemplo.com/pkg/domain/entities"
)

func TestAccountUseCase_CreateAccount(t *testing.T) {
	t.Run("should successfully create an account and return it", func(t *testing.T) {

		storage := make(map[int]entities.Account)
		accountUsecase := NewAccountUseCase(0, storage)
		account := entities.Account{Name: "John Doe", Cpf: "11111111030", Balance: 0}

		createdAccount, err := accountUsecase.CreateAccount(account)

		if createdAccount != account {
			t.Errorf("expected %+v but got %+v", account, createdAccount)
		}

		if err != nil {
			t.Errorf("expected nil err but got %v", err)
		}
	})

	t.Run("should return err when trying to create account with already created cpf", func(t *testing.T) {

		storage := make(map[int]entities.Account)
		currentAccount := entities.Account{Id: 0, Name: "John Doe", Cpf: "11111111030"}
		storage[0] = currentAccount
		accountUsecase := NewAccountUseCase(0, storage)

		createdAccount, err := accountUsecase.CreateAccount(currentAccount)

		if createdAccount != (entities.Account{}) {
			t.Errorf("expected blank account, but got %+v", createdAccount)
		}

		if err == nil {
			t.Error("expected err but got nil")
		}
	})
}

func TestAccountUseCase_GetBalanceById(t *testing.T) {

	t.Run("should return an account when id is found", func(t *testing.T) {

		account := entities.Account{Id: 0, Name: "John Doe", Cpf: "11111111030", Balance: 10.0}
		storage := make(map[int]entities.Account)
		AccountUseCase := NewAccountUseCase(0, storage)
		storage[0] = account

		getBalance, err := AccountUseCase.GetBalanceById(account.Id)

		if getBalance == 0 {
			t.Errorf("Balance account %g should be different from 0", getBalance)
		}

		if err != nil {
			t.Errorf("Err should be nil but it is %s", err)
		}

	})

	t.Run("should return a blank account when id isn't found", func(t *testing.T) {

		storage := make(map[int]entities.Account)
		AccountUseCase := NewAccountUseCase(0, storage)

		getBalance, err := AccountUseCase.GetBalanceById(1)

		if getBalance != 0 {
			t.Errorf("Balance Account should be 0 but it is %g", getBalance)
		}

		if err == nil {
			t.Errorf("Err should be different from nil but it is %s", err)
		}

	})

}

func TestAccountUseCase_GetAccounts(t *testing.T) {

	t.Run("should return a full list of accounts", func(t *testing.T) {

		account := entities.Account{Id: 0, Name: "John Doe", Cpf: "11111111030"}
		storage := make(map[int]entities.Account)
		AccountUseCase := NewAccountUseCase(0, storage)
		storage[0] = account

		getAccounts := AccountUseCase.GetAccounts()

		if len(getAccounts) == 0 {
			t.Error("expected a full list")
		}

	})

	t.Run("should return an empty list", func(t *testing.T) {

		storage := make(map[int]entities.Account)
		AccountUseCase := NewAccountUseCase(0, storage)

		getAccounts := AccountUseCase.GetAccounts()

		if len(getAccounts) != 0 {
			t.Error("expected a empty list")
		}

	})

}

func TestAccountUseCase_CheckAccounts(t *testing.T) {

	t.Run("should return nil when accounts have already been created", func(t *testing.T) {

		account := entities.Account{Id: 0, Name: "John Doe", Cpf: "11111111030"}
		storage := make(map[int]entities.Account)
		AccountUseCase := NewAccountUseCase(0, storage)
		storage[0] = account
		err := AccountUseCase.CheckAccounts(account.Id)

		if err != nil {
			t.Error("expected nil when account exists")
		}

	})

	t.Run("should return an err message when id isn't found", func(t *testing.T) {

		storage := make(map[int]entities.Account)
		AccountUseCase := NewAccountUseCase(0, storage)
		err := AccountUseCase.CheckAccounts(0)

		if err == nil {
			t.Error("expected a err message")
		}

	})

}

func TestAccountUseCase_UpdateAccountBalance(t *testing.T) {

	t.Run("Should return nil when account was updated", func(t *testing.T) {

		account := entities.Account{Id: 0, Name: "John Doe", Cpf: "11111111030", Balance: 0}
		storage := make(map[int]entities.Account)
		AccountUseCase := NewAccountUseCase(0, storage)
		storage[0] = account

		ErrUpdateAccount := AccountUseCase.UpdateAccountBalance(account.Id, 20.0)

		if ErrUpdateAccount != nil {
			t.Errorf("Expected nil but got %s", ErrUpdateAccount)
		}

	})

	t.Run("Should return an err massage when account don't exists", func(t *testing.T) {

		storage := make(map[int]entities.Account)
		AccountUseCase := NewAccountUseCase(0, storage)

		ErrUpdateAccount := AccountUseCase.UpdateAccountBalance(0, 20.0)

		if ErrUpdateAccount == nil {
			t.Error("Expected err message but got nil")
		}

	})

	t.Run("Should return an err message when balance account is less than zero", func(t *testing.T) {

		account := entities.Account{Id: 0, Name: "John Doe", Cpf: "11111111030", Balance: -10}
		storage := make(map[int]entities.Account)
		AccountUseCase := NewAccountUseCase(0, storage)
		storage[0] = account

		ErrUpdateAccount := AccountUseCase.UpdateAccountBalance(account.Id, account.Balance)

		if ErrUpdateAccount == nil {
			t.Error("Expected err message")
		}

	})
}

func TestAccountUseCase_GetAccountById(t *testing.T) {

	t.Run("Should return an account when the searched account is found", func(t *testing.T) {

		account := entities.Account{Id: 0, Name: "John Doe", Cpf: "11111111030"}
		storage := make(map[int]entities.Account)
		AccountUseCase := NewAccountUseCase(0, storage)
		storage[0] = account

		GetAccountById, err := AccountUseCase.GetAccountByID(account.Id)

		if GetAccountById == (entities.Account{}) {
			t.Errorf("Expected account but got %+v", GetAccountById)
		}

		if err != nil {
			t.Error("Expected err equal nil")
		}

	})

	t.Run("Should return an empty account and a err message when account don't exists", func(t *testing.T) {

		storage := make(map[int]entities.Account)
		AccountUseCase := NewAccountUseCase(0, storage)

		GetAccountById, err := AccountUseCase.GetAccountByID(0)

		if GetAccountById != (entities.Account{}) {
			t.Errorf("Expected empty account but got %+v", GetAccountById)
		}

		if err == nil {
			t.Error("Expected err different from nil")
		}

	})

}
