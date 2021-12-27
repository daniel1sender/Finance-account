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

func TestAccountUseCase_Create(t *testing.T) {

 	accountFile := "Account_Repository.json"
	openAccountFile, err := os.OpenFile(accountFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error to open file: %v", err)
	} 

	t.Run("should successfully create an account and return it", func(t *testing.T) {
		//_ = os.Remove("Account_Repository.json")
		storageFiles := accounts_repository.NewStorage(openAccountFile)
		accountUsecase := NewUseCase(storageFiles)

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		createdAccount, err := accountUsecase.Create(name, cpf, secret, balance)

		if err != nil {
			t.Errorf("expected no error but got '%s'", err)
		}

		if createdAccount == (entities.Account{}) {
			t.Errorf("expected an account but got %+v", createdAccount)
		}

	})

	t.Run("should return error when trying to create account with already created cpf account", func(t *testing.T) {
		//_ = os.Remove("Account_Repository.json")
		storageFiles := accounts_repository.NewStorage(openAccountFile)
		accountUsecase := NewUseCase(storageFiles)

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
