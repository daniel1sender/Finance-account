package accounts

import (
	"errors"
	"testing"

	"exemplo.com/pkg/domain/entities"
)

func TestAccountUseCase_CreateAccount(t *testing.T) {
	t.Run("should successfully create an account and return it", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		accountUsecase := NewAccountUseCase(storage)
		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		createdAccount, err := accountUsecase.CreateAccount(name, cpf, secret, balance)

		if err != nil {
			t.Error("Expected nil error but got %w", err)
		}

		if createdAccount == (entities.Account{}) {
			t.Errorf("Expected an account but got %v", createdAccount)
		}

	})

	t.Run("should return error when trying to create account with already created cpf account", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		accountUsecase := NewAccountUseCase(storage)

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		createdAccount, err := accountUsecase.CreateAccount(name, cpf, secret, balance)

		if err != nil {
			t.Errorf("Expected nil error but got %s", err)
		}

		if createdAccount == (entities.Account{}) {
			t.Errorf("Expected %+v but got %+v", entities.Account{}, createdAccount)
		}

		createdAccount1, err1 := accountUsecase.CreateAccount(name, cpf, secret, balance)

		if !errors.Is(err1, ErrExistingCPF) {
			t.Errorf("Expected %s but got %s", ErrExistingCPF, err1)
		}

		if createdAccount1 != (entities.Account{}) {
			t.Errorf("Expected %+v but got %+v", entities.Account{}, createdAccount1)
		}

	})
}

func TestAccountUseCase_GetBalanceByID(t *testing.T) {

	t.Run("should return an account when id is found", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		AccountUseCase := NewAccountUseCase(storage)
		account, err := entities.NewAccount("John Doe", "11111111030", "123", 10)
		if err != nil {
			t.Error("Err should be nil if account was successfully created")
		}
		storage[account.ID] = account

		getBalance, err := AccountUseCase.GetBalanceByID(account.ID)

		if getBalance == 0 {
			t.Errorf("Balance account %d should be different from 0", getBalance)
		}

		if err != nil {
			t.Errorf("Err should be nil but it is %s", err)
		}

	})

	t.Run("should return a blank account when id isn't found", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		AccountUseCase := NewAccountUseCase(storage)
		account, err := entities.NewAccount("John Doe", "11111111030", "123", 10)
		if err != nil {
			t.Error("error should be nil if account was successfully created")
		}

		getBalance, err := AccountUseCase.GetBalanceByID(account.ID)

		if getBalance != 0 {
			t.Errorf("balance Account should be 0 but it is %d", getBalance)
		}

		if err == nil {
			t.Errorf("error should be different from nil but it is %s", err)
		}

	})

}

func TestAccountUseCase_GetAccounts(t *testing.T) {

	t.Run("should return a full list of accounts", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		AccountUseCase := NewAccountUseCase(storage)
		account, err := entities.NewAccount("John Doe", "11111111030", "123", 10)
		if err != nil {
			t.Error("error should be nil if account was successfully created")
		}
		storage[account.ID] = account

		getAccounts := AccountUseCase.GetAccounts()

		if len(getAccounts) == 0 {
			t.Error("expected a full list of accounts")
		}

	})

	t.Run("should return an empty list", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		AccountUseCase := NewAccountUseCase(storage)

		getAccounts := AccountUseCase.GetAccounts()

		if len(getAccounts) != 0 {
			t.Error("expected an empty list")
		}

	})

}

func TestAccountUseCase_CheckAccounts(t *testing.T) {

	t.Run("should return nil when accounts have already been created", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		AccountUseCase := NewAccountUseCase(storage)
		account, err := entities.NewAccount("John Doe", "11111111030", "123", 10)
		if err != nil {
			t.Error("error should be nil if account was successfully created")
		}

		storage[account.ID] = account

		CheckAccountsError := AccountUseCase.CheckAccounts(account.ID)

		if CheckAccountsError != nil {
			t.Error("expected nil when account exists")
		}

	})

	t.Run("should return an error message when id isn't found", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		AccountUseCase := NewAccountUseCase(storage)
		account, err := entities.NewAccount("John Doe", "11111111030", "123", 10)
		if err != nil {
			t.Error("error should be nil if account was successfully created")
		}

		CheckAccountsError := AccountUseCase.CheckAccounts(account.ID)

		if CheckAccountsError == nil {
			t.Error("expected a error message")
		}

	})

}

func TestAccountUseCase_UpdateAccountBalance(t *testing.T) {

	t.Run("Should return nil when account was updated", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		AccountUseCase := NewAccountUseCase(storage)
		account, err := entities.NewAccount("John Doe", "11111111030", "123", 10)
		if err != nil {
			t.Error("error should be nil if account was successfully created")
		}
		storage[account.ID] = account

		UpdateAccountError := AccountUseCase.UpdateAccountBalance(account.ID, 20.0)

		if UpdateAccountError != nil {
			t.Errorf("Expected nil but got %s", UpdateAccountError)
		}

	})

	t.Run("Should return an error massage when account don't exists", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		AccountUseCase := NewAccountUseCase(storage)

		//passando qualquer id, sem criar a conta
		err := AccountUseCase.UpdateAccountBalance("1", 20.0)

		if err != ErrIDNotFound {
			t.Errorf("Expected %s but got %s", ErrIDNotFound, err)
		}

	})

	t.Run("Should return an error message when balance account is less than zero", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		AccountUseCase := NewAccountUseCase(storage)

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("Expected nil error to create a new account butgot %s", err)
		}
		storage[account.ID] = account

		err = AccountUseCase.UpdateAccountBalance(account.ID, -10)

		if !errors.Is(err, ErrBalanceLessThanZero) {
			t.Errorf("Expected %s but got %s", ErrBalanceLessThanZero, err)
		}

	})
}

func TestAccountUseCase_GetAccountById(t *testing.T) {

	t.Run("Should return an account when the searched account is found", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		AccountUseCase := NewAccountUseCase(storage)

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("Expected nil error to create a new account but got %s", err)
		}
		storage[account.ID] = account
		GetAccountByID, err := AccountUseCase.GetAccountByID(account.ID)

		if GetAccountByID == (entities.Account{}) {
			t.Errorf("Expected account but got %+v", GetAccountByID)
		}

		if err != nil {
			t.Error("Expected error equal nil")
		}

	})

	t.Run("Should return an empty account and a error message when account don't exist", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		AccountUseCase := NewAccountUseCase(storage)

		//passando qualquer id
		GetAccountByID, err := AccountUseCase.GetAccountByID("account.ID")

		if GetAccountByID != (entities.Account{}) {
			t.Errorf("Expected empty account but got %+v", GetAccountByID)
		}

		if !errors.Is(err, ErrIDNotFound) {
			t.Errorf("Expected %s but got %s", ErrIDNotFound, err)
		}

	})

}
