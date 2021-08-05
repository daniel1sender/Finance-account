package accounts

import (
	"testing"
	"time"

	"exemplo.com/pkg/domain/entities"
)

func TestAccountUseCase_CreateAccount(t *testing.T) {
	t.Run("should successfully create an account and return it", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		accountUsecase := NewAccountUseCase(storage)
		account, err := entities.NewAccount("John Doe", "11111111030", "123", 10)

		if err != nil {
			t.Error("Err should be nil if account was successfully created")
		}

		createdAccount, err := accountUsecase.CreateAccount("John Doe", "11111111030", "123", 10)

		if createdAccount == (entities.Account{}) {
			t.Error("returned account is empty")
		}

		if createdAccount.CreatedAt == (time.Time{}) {
			t.Error("returned createAt is empty")
		}

		if createdAccount.Secret == account.Secret {
			t.Errorf("expected %+v but got %+v", account.Secret, createdAccount.Secret)
		}

		if err != nil {
			t.Errorf("expected nil err but got %v", err)
		}

		//checar se o tempo não é zero(time.Time{}), e se não retorna uma conta fazia
	})

	t.Run("should return err when trying to create account with already created account cpf", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		accountUsecase := NewAccountUseCase(storage)
		account, err := entities.NewAccount("John Doe", "11111111030", "123", 10)
		if err != nil {
			t.Error("Err should be nil if account was successfully created")
		}
		storage[account.Id] = account

		createdAccount, err := accountUsecase.CreateAccount("John Doe", "11111111030", "123", 10)

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

		storage := make(map[string]entities.Account)
		AccountUseCase := NewAccountUseCase(storage)
		account, err := entities.NewAccount("John Doe", "11111111030", "123", 10)
		if err != nil {
			t.Error("Err should be nil if account was successfully created")
		}
		storage[account.Id] = account

		getBalance, err := AccountUseCase.GetBalanceById(account.Id)

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

		getBalance, err := AccountUseCase.GetBalanceById(" ")

		if getBalance != 0 {
			t.Errorf("Balance Account should be 0 but it is %d", getBalance)
		}

		if err == nil {
			t.Errorf("Err should be different from nil but it is %s", err)
		}

	})

}

func TestAccountUseCase_GetAccounts(t *testing.T) {

	t.Run("should return a full list of accounts", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		AccountUseCase := NewAccountUseCase(storage)
		account, err := entities.NewAccount("John Doe", "11111111030", "123", 10)
		if err != nil {
			t.Error("Err should be nil if account was successfully created")
		}
		storage[account.Id] = account

		getAccounts := AccountUseCase.GetAccounts()

		if len(getAccounts) == 0 {
			t.Error("expected a full list")
		}

	})

	t.Run("should return an empty list", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		AccountUseCase := NewAccountUseCase(storage)

		getAccounts := AccountUseCase.GetAccounts()

		if len(getAccounts) != 0 {
			t.Error("expected a empty list")
		}

	})

}

func TestAccountUseCase_CheckAccounts(t *testing.T) {

	t.Run("should return nil when accounts have already been created", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		AccountUseCase := NewAccountUseCase(storage)
		account, err := entities.NewAccount("John Doe", "11111111030", "123", 10)
		if err != nil {
			t.Error("err should be nil if account was successfully created")
		}

		storage[account.Id] = account

		CheckAccountsError := AccountUseCase.CheckAccounts(account.Id)

		if CheckAccountsError != nil {
			t.Error("expected nil when account exists")
		}

	})

	t.Run("should return an err message when id isn't found", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		AccountUseCase := NewAccountUseCase(storage)
		account, err := entities.NewAccount("John Doe", "11111111030", "123", 10)
		if err != nil {
			t.Error("err should be nil if account was successfully created")
		}

		CheckAccountsError := AccountUseCase.CheckAccounts(account.Id)

		if CheckAccountsError == nil {
			t.Error("expected a err message")
		}

	})

}

func TestAccountUseCase_UpdateAccountBalance(t *testing.T) {

	t.Run("Should return nil when account was updated", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		AccountUseCase := NewAccountUseCase(storage)
		account, err := entities.NewAccount("John Doe", "11111111030", "123", 10)
		if err != nil {
			t.Error("err should be nil if account was successfully created")
		}
		storage[account.Id] = account

		UpdateAccountError := AccountUseCase.UpdateAccountBalance(account.Id, 20.0)

		if UpdateAccountError != nil {
			t.Errorf("Expected nil but got %s", UpdateAccountError)
		}

	})

	t.Run("Should return an err massage when account don't exists", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		AccountUseCase := NewAccountUseCase(storage)
		account, err := entities.NewAccount("John Doe", "11111111030", "123", 10)
		if err != nil {
			t.Error("err should be nil if account was successfully created")
		}
		ErrUpdateAccount := AccountUseCase.UpdateAccountBalance(account.Id, 20.0)

		if ErrUpdateAccount == nil {
			t.Error("Expected err message but got nil")
		}

	})

	t.Run("Should return an err message when balance account is less than zero", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		AccountUseCase := NewAccountUseCase(storage)
		account, err := entities.NewAccount("John Doe", "11111111030", "123", -10)
		if err != nil {
			t.Error("err should be nil if account was successfully created")
		}
		storage[account.Id] = account

		ErrUpdateAccount := AccountUseCase.UpdateAccountBalance(account.Id, account.Balance)

		if ErrUpdateAccount == nil {
			t.Error("Expected err message")
		}

	})
}

func TestAccountUseCase_GetAccountById(t *testing.T) {

	t.Run("Should return an account when the searched account is found", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		AccountUseCase := NewAccountUseCase(storage)
		account, err := entities.NewAccount("John Doe", "11111111030", "123", -10)
		if err != nil {
			t.Error("err should be nil if account was successfully created")
		}
		storage[account.Id] = account

		GetAccountById, err := AccountUseCase.GetAccountByID(account.Id)

		if GetAccountById == (entities.Account{}) {
			t.Errorf("Expected account but got %+v", GetAccountById)
		}

		if err != nil {
			t.Error("Expected err equal nil")
		}

	})

	t.Run("Should return an empty account and a err message when account don't exists", func(t *testing.T) {

		storage := make(map[string]entities.Account)
		AccountUseCase := NewAccountUseCase(storage)
		account, err := entities.NewAccount("John Doe", "11111111030", "123", -10)
		if err != nil {
			t.Error("err should be nil if account was successfully created")
		}
		GetAccountById, err := AccountUseCase.GetAccountByID(account.Id)

		if GetAccountById != (entities.Account{}) {
			t.Errorf("Expected empty account but got %+v", GetAccountById)
		}

		if err == nil {
			t.Error("Expected err different from nil")
		}

	})

}
