package domain_test

import (
	"fmt"
	"testing"

	"exemplo.com/domain"
)

func TestCreateAccount(t *testing.T) {
	//Testar se o cpf passado tem 11 elementos
	//Caso o cpf não tenha 11 digitos a conta não é criada
	//Caso o cpf tenha 11 digitos a conta é criada

	t.Run("If the Cpf does not have 11 digits the account is not created", func(t *testing.T) {
		// preparação
		account := domain.Account{
			Id:      1,
			Name:    "daniel",
			Cpf:     "1333330",
			Balance: 15,
		}
		expectedErrMessage := fmt.Sprintf("CPF %s is not correct", account.Cpf)

		// teste
		accountMap, err := domain.CreateAccount(account)

		//assert(verificação)
		if accountMap != nil {
			t.Errorf("accountMap must be nil and it is %v", accountMap)
		}

		if err.Error() != expectedErrMessage {
			t.Errorf("err must be %s but it is %s", expectedErrMessage, err)
		}
	})

	t.Run("If the Cpf has 11 digits the account is created", func(t *testing.T) {
		account := domain.Account{
			Id:      2,
			Name:    "joão",
			Cpf:     "12345678910",
			Balance: 20,
		}

		expectedErrMessage := fmt.Sprintf("CPF %s is not correct", account.Cpf)

		accountMap, err := domain.CreateAccount(account)

		if accountMap == nil {
			t.Errorf("accountMap must not be nil and it is %v", accountMap)
		}

		if err != nil && err.Error() == expectedErrMessage {
			t.Errorf("err must be %s but it is %s", err, expectedErrMessage)
		}
	})
}

func TestGetBalanceById(t *testing.T) {
	t.Run("if the id is not found", func(t *testing.T) {

		id := 5
		expectedErrMessage := fmt.Sprintf("no id %d found", id)

		BalanceAccount, err := domain.GetBalanceById(id)

		if BalanceAccount != 0 {
			t.Errorf("BalanceAccount must be nill but it is %d", id)
		}

		if err.Error() != expectedErrMessage {
			t.Errorf("err must be %s but it is %s", expectedErrMessage, err)
		}

	})

	t.Run("if the id it is found", func(t *testing.T) {

		id := 0

		fmt.Printf("A account was created %v\n", domain.AccountsMap)

		//expectedErrMessage := fmt.Sprintf("no id %d found", id)

		BalanceAccount, err := domain.GetBalanceById(id)

		if BalanceAccount == 0 {
			t.Errorf("BalanceAccount must be %g but it is nil", BalanceAccount)
		}

		if err != nil {
			t.Errorf("err must be nil")
		}
	})
}

func TestGetAccounts(t *testing.T) {
	//limpar o map
	//numero de contas estar zero
	//testar com o tamanho da lista

	t.Run("If accounts was created", func(t *testing.T) {

		AccountsList := domain.GetAccounts()

		if len(AccountsList) == 0 {
			t.Errorf("Accounts created")
		}

	})

	t.Run("should return a empty list when no account was created", func(t *testing.T) {

		domain.AccountsMap = make(map[int]domain.Account)

		fmt.Println(domain.AccountsMap)

		AccountsList := domain.GetAccounts()

		if len(AccountsList) != 0 {
			t.Errorf("No account created")
		}
	})
}

func TestCheckAccounts(t *testing.T) {

	t.Run("should return nil when accounts exist", func(t *testing.T) {

		account := domain.Account{
			Id:      2,
			Name:    "joão",
			Cpf:     "12345678910",
			Balance: 20,
		}
		domain.AccountsMap[0] = account
		id := 0
		result := domain.CheckAccounts(id)

		if result != nil {
			t.Errorf("result should be nil but it is errMessage")
		}
	})

	t.Run("should return nil when accounts exist", func(t *testing.T) {
		id := 5
		result := domain.CheckAccounts(id)

		if result == nil {
			t.Errorf("result should be err message but it is nil")
		}
	})
}

func TestAccountTransfer(t *testing.T) {
	t.Run("Should return two maps when the transfer is made", func(t *testing.T) {

		account1 := domain.Account{
			Id:      2,
			Name:    "joão",
			Cpf:     "12345678910",
			Balance: 10,
		}

		account2 := domain.Account{
			Id:      2,
			Name:    "joão",
			Cpf:     "12345678910",
			Balance: 1,
		}

		domain.AccountsMap[0] = account1
		domain.AccountsMap[1] = account2

		transfer := domain.Transfer{"1", 0, 1, 1}

		originEmptyAccount, destinationEmptyAccount := domain.Account{}, domain.Account{}

		originAccount, destinationAccount, err := domain.AccountTransfer(transfer)

		expectedErrMessage := "the origin account does not have enough balance"

		if originAccount == originEmptyAccount {
			t.Errorf("struct of origin account returned %v should not be empty", originAccount)
		}

		if destinationAccount == destinationEmptyAccount {
			t.Errorf("struct of destination account returned %v should not be empty", destinationAccount)
		}

		if err != nil && err.Error() == expectedErrMessage {
			t.Errorf("err message %s should be nil", err)
		}

	})
}
