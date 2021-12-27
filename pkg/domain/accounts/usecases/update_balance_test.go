package usecases

import (
	"errors"
	"log"
	"os"
	"testing"

	accounts_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	accounts_repository "github.com/daniel1sender/Desafio-API/pkg/gateways/store/files/accounts"
)

func TestAccountUseCase_UpdateBalance(t *testing.T) {
	accountFile := "Account_Repository.json"
	openAccountFile, err := os.OpenFile(accountFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error to open file: %v", err)
	}
	t.Run("should return nil when account was updated", func(t *testing.T) {

		storageFiles := accounts_repository.NewStorage(openAccountFile)
		accountUseCase := NewUseCase(storageFiles)
		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}

		storageFiles.Upsert(account)

		updateAccountError := accountUseCase.UpdateBalance(account.ID, 20.0)

		if updateAccountError != nil {
			t.Errorf("expected no error but got '%s'", updateAccountError)
		}

	})

	t.Run("should return an error massage when account don't exists", func(t *testing.T) {
		storageFiles := accounts_repository.NewStorage(openAccountFile)
		accountUseCase := NewUseCase(storageFiles)

		err := accountUseCase.UpdateBalance("1", 20.0)

		if !errors.Is(err, accounts_usecase.ErrIDNotFound) {
			t.Errorf("expected '%s' but got '%s'", accounts_usecase.ErrIDNotFound, err)
		}

	})

	t.Run("should return an error message when balance account is less than zero", func(t *testing.T) {

		storageFiles := accounts_repository.NewStorage(openAccountFile)
		accountUseCase := NewUseCase(storageFiles)

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}

		storageFiles.Upsert(account)

		err = accountUseCase.UpdateBalance(account.ID, -10)

		if !errors.Is(err, ErrBalanceLessZero) {
			t.Errorf("expected '%s' but got '%s'", ErrBalanceLessZero, err)
		}

	})
}
