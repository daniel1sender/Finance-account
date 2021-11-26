package usecases

import (
	"errors"
	"os"
	"testing"

	accounts_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	accounts_repository "github.com/daniel1sender/Desafio-API/pkg/gateways/store/repository/accounts"
)

func TestAccountUseCase_UpdateBalance(t *testing.T) {

	t.Run("should return nil when account was updated", func(t *testing.T) {

		//storage := accounts.NewStorage()
		//accountUseCase := NewUseCase(storage)
		storageFiles := accounts_repository.NewStorage()
		accountUseCase := NewUseCase(storageFiles)
		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}

		//storage.Upsert(account)
		storageFiles.Upsert(account)

		updateAccountError := accountUseCase.UpdateBalance(account.ID, 20.0)

		if updateAccountError != nil {
			t.Errorf("expected no error but got '%s'", updateAccountError)
		}

	})

	t.Run("should return an error massage when account don't exists", func(t *testing.T) {
		_ = os.Remove("Account_Repository.json")
		//storage := accounts.NewStorage()
		//accountUseCase := NewUseCase(storage)
		storageFiles := accounts_repository.NewStorage()
		accountUseCase := NewUseCase(storageFiles)

		//passando qualquer id, sem criar a conta
		err := accountUseCase.UpdateBalance("1", 20.0)

		if err != accounts_usecase.ErrIDNotFound {
			t.Errorf("expected '%s' but got '%s'", accounts_usecase.ErrIDNotFound, err)
		}

	})

	t.Run("should return an error message when balance account is less than zero", func(t *testing.T) {

		//storage := accounts.NewStorage()
		//accountUseCase := NewUseCase(storage)
		storageFiles := accounts_repository.NewStorage()
		accountUseCase := NewUseCase(storageFiles)

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := entities.NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("expected no error to create a new account but got '%s'", err)
		}

		//storage.Upsert(account)
		storageFiles.Upsert(account)

		err = accountUseCase.UpdateBalance(account.ID, -10)

		if !errors.Is(err, ErrBalanceLessZero) {
			t.Errorf("expected '%s' but got '%s'", ErrBalanceLessZero, err)
		}

	})
}
