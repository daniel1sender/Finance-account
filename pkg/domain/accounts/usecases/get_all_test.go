package usecases

import (
	"os"
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	accounts_repository "github.com/daniel1sender/Desafio-API/pkg/gateways/store/repository/accounts"
)

func TestAccountUseCase_Get(t *testing.T) {

	t.Run("should return a full list of accounts", func(t *testing.T) {

		//storageMemory := accounts.NewStorage()
		//accountUseCase := NewUseCase(storageMemory)
		storageFiles := accounts_repository.NewStorage()
		accountUsecase := NewUseCase(storageFiles)
		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}

		//storageMemory.Upsert(account)
		storageFiles.Upsert(account)

		getAccounts := accountUsecase.GetAll()

		if len(getAccounts) == 0 {
			t.Error("expected a full list of accounts")
		}

	})

	t.Run("should return an empty list", func(t *testing.T) {
		_ = os.Remove("Account_Repository.json")
		storageFiles := accounts_repository.NewStorage()
		accountUsecase := NewUseCase(storageFiles)

		getAccounts := accountUsecase.GetAll()

		if len(getAccounts) != 0 {
			t.Error("expected an empty list")
		}

	})

}
