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

func TestAccountUseCase_GetById(t *testing.T) {
	accountFile := "Account_Repository.json"
	openAccountFile, err := os.OpenFile(accountFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error to open file: %v", err)
	}

	t.Run("should return an account when the searched account is found", func(t *testing.T) {

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
		getAccountByID, err := accountUseCase.GetByID(account.ID)

		if getAccountByID == (entities.Account{}) {
			t.Errorf("expected an account but got %+v", getAccountByID)
		}

		if err != nil {
			t.Errorf("expected error equal nil but got '%s'", err)
		}

	})

	t.Run("should return an empty account and a error message when account don't exist", func(t *testing.T) {

		_ = os.Remove("Account_Repository.json")
		storageFiles := accounts_repository.NewStorage(openAccountFile)
		accountUseCase := NewUseCase(storageFiles)

		getAccountByID, err := accountUseCase.GetByID("account.ID")

		if getAccountByID != (entities.Account{}) {
			t.Errorf("expected empty account but got %+v", getAccountByID)
		}

		if !errors.Is(err, accounts_usecase.ErrIDNotFound) {
			t.Errorf("expected '%s' but got '%s'", accounts_usecase.ErrIDNotFound, err)
		}

	})

}
