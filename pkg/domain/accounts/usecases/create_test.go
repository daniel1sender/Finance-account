package usecases

import (
	"errors"
	"os"
	"testing"

	accounts_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/memory/accounts"
	accounts_repository "github.com/daniel1sender/Desafio-API/pkg/gateways/store/repository/accounts"
)

func TestAccountUseCase_Create(t *testing.T) {

	t.Run("should successfully create an account and return it", func(t *testing.T) {
		_ = os.Remove("Account_Repository.json")
		//storageMemory := accounts.NewStorage()
		//accountUsecase := NewUseCase(storageMemory)
		storageFiles := accounts_repository.NewStorage()
		accountUsecase := NewUseCase(storageFiles)

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account := entities.Account{Name: name, CPF: cpf, Secret: secret, Balance: balance}

		createdAccount, err := accountUsecase.Create(account.Name, account.CPF, account.Secret, account.Balance)

		if createdAccount == (entities.Account{}) {
			t.Errorf("expected an account but got %+v", createdAccount)
		}

		if createdAccount.Name != account.Name {
			t.Errorf("expected '%s' but got '%s'", account.Name, createdAccount.Name)
		}

		if createdAccount.CPF != account.CPF {
			t.Errorf("expected '%s' but got '%s'", account.CPF, createdAccount.CPF)
		}

		if createdAccount.Balance != account.Balance {
			t.Errorf("expected '%d' but got '%d'", account.Balance, createdAccount.Balance)
		}

		if err != nil {
			t.Errorf("expected no error but got '%s'", err)
		}

	})

	t.Run("should return an empty account and an error when trying to create account with already created cpf account", func(t *testing.T) {

		storageMemory := accounts.NewStorage()
		accountUsecase := NewUseCase(storageMemory)

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		createdAccount, err := accountUsecase.Create(name, cpf, secret, balance)

		if err != nil {
			t.Errorf("expected no error but got '%s'", err)
		}

		if createdAccount == (entities.Account{}) {
			t.Errorf("expected %+v but got %+v", entities.Account{}, createdAccount)
		}

		createdAccount1, err1 := accountUsecase.Create(name, cpf, secret, balance)

		if !errors.Is(err1, accounts_usecase.ErrExistingCPF) {
			t.Errorf("expected '%s' but got '%s'", accounts_usecase.ErrExistingCPF, err1)
		}

		if createdAccount1 != (entities.Account{}) {
			t.Errorf("expected %+v but got %+v", entities.Account{}, createdAccount1)
		}

	})
}
