package accounts

import (
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func TestGetAll(t *testing.T) {

	t.Run("should return a list of accounts when accounts have already been created", func(t *testing.T) {
		storage := NewStorage()
		account := entities.Account{ID: "123", Balance: 10}
		storage.Users[account.ID] = account
		returnedAccount := storage.GetAll()

		if len(returnedAccount) == 0 {
			t.Errorf("expected a full list of accounts")
		}
	})

	t.Run("should return a list of accounts when accounts don't have already been created", func(t *testing.T) {
		storage := NewStorage()
		returnedAccount := storage.GetAll()

		if len(returnedAccount) != 0 {
			t.Errorf("expected an empty list of accounts")
		}
	})

}
