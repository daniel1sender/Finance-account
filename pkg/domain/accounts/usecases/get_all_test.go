package usecases

import (
	"log"
	"os"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	accounts_repository "github.com/daniel1sender/Desafio-API/pkg/gateways/store/files/accounts"
)

func TestAccountUseCase_Get(t *testing.T) {
	accountFile := "Account_Repository.json"
	openAccountFile, err := os.OpenFile(accountFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error to open file: %v", err)
	}
	t.Run("should return a full list of accounts", func(t *testing.T) {

		storageFiles := accounts_repository.NewStorage(openAccountFile)
		accountUsecase := NewUseCase(storageFiles)
		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}

		storageFiles.Upsert(account)

		getAccounts := accountUsecase.GetAll()

		if len(getAccounts) == 0 {
			t.Error("expected a full list of accounts")
		}

	})

	t.Run("should return an empty list", func(t *testing.T) {
		_ = os.Remove("Account_Repository.json")
		storageFiles := accounts_repository.NewStorage(openAccountFile)
		accountUsecase := NewUseCase(storageFiles)

		getAccounts := accountUsecase.GetAll()

		if len(getAccounts) != 0 {
			t.Error("expected an empty list")
		}

	})

}
