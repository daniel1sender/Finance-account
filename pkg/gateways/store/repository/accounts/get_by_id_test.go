package accounts

import (
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func TestGetByID(t *testing.T) {

	t.Run("should return an account when id informed already exists", func(t *testing.T) {
		storage := NewStorage()
		account := entities.Account{ID: "123", Balance: 10}
		storage.users[account.ID] = account
		returnedAccount, err := storage.GetByID(account.ID)
		if err != nil {
			t.Errorf("expected null error but got %v", err)
		}
		if returnedAccount != account {
			t.Errorf("expected %v but got %v", account, returnedAccount)
		}
	})

	t.Run("should return an error when account it is not found by id", func(t *testing.T) {
		storage := NewStorage()
		account := entities.Account{ID: "123", Balance: 10}
		returnedAccount, err := storage.GetByID(account.ID)
		if err == nil {
			t.Errorf("expected null error but got '%v'", err)
		}
		if returnedAccount != (entities.Account{}) {
			t.Errorf("expected %v but got %v", entities.Account{}, returnedAccount)
		}
	})

}
