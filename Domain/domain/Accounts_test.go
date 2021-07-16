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
			Id:      "1",
			Name:    "daniel",
			Cpf:     "1333330",
			Balance: 0,
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
			Id:      "2",
			Name:    "joão",
			Cpf:     "12345678910",
			Balance: 0,
		}

		expectedErrMessage := fmt.Sprintf("CPF %s is not correct", account.Cpf)

		accountMap, err := domain.CreateAccount(account)

		if accountMap == nil {
			t.Errorf("accountMap must not be nil and it is %v", accountMap)
		}

		if err.Error() == expectedErrMessage {
			t.Errorf("err must be %s but it is %s", err, expectedErrMessage)
		}
	})
}

func TestGetAccounts(t *testing.T) {
	t.Run("If any account was created", func(t *testing.T) {

		AccountsList := domain.GetAccounts()

		if AccountsList != nil {
			t.Errorf("No account created")
		}
	})

	t.Run("If accounts was created", func(t *testing.T) {

		AccountsList := domain.GetAccounts()

		if AccountsList == nil {
			t.Errorf("Accounts created")
		}

	})
}

func TestGetBalanceById(t *testing.T) {
	t.Run("if the id is not found", func(t *testing.T) {

		id := "5"
		expectedErrMessage := fmt.Sprintf("No id %s found", id)

		BalanceAccount, err := domain.GetBalanceById(id)

		if BalanceAccount != 0 {
			t.Errorf("BalanceAccount must be nill but it is %s", id)
		}

		if err.Error() != expectedErrMessage {
			t.Errorf("err must be %s but it is %s", expectedErrMessage, err)
		}

	})

	t.Run("if the id it is found", func(t *testing.T) {

		id := "3"

		expectedErrMessage := fmt.Sprintf("No id %s found", id)

		BalanceAccount, err := domain.GetBalanceById(id)

		if BalanceAccount == 0 {
			t.Errorf("BalanceAccount must be %d but it is nil", BalanceAccount)
		}

		if err.Error() == expectedErrMessage {
			t.Errorf("err must be nil but it is %s", expectedErrMessage)
		}
	})
}
