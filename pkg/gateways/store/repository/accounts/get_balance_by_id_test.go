package accounts

import (
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func TestGetBalanceByID(t *testing.T) {
	t.Run("should return the balance when account it is found by id", func(t *testing.T) {
		storage := NewStorage()
		account := entities.Account{ID: "123", Balance: 10}
		storage.Users[account.ID] = account
		balance, err := storage.GetBalanceByID(account.ID)
		if err != nil {
			t.Errorf("expected nil error but got '%v'", err)
		}
		if balance != account.Balance {
			t.Errorf("expected '%d' but got '%d'", account.Balance, balance)
		}
	})

	t.Run("should return an error when account it is not found by id", func(t *testing.T) {
		storage := NewStorage()
		account := entities.Account{ID: "123", Balance: 10}
		balance, err := storage.GetBalanceByID(account.ID)
		if err == nil {
			t.Errorf("expected nil error but got '%v'", err)
		}
		if balance == account.Balance {
			t.Errorf("expected '%d' but got '%d'", account.Balance, balance)
		}
	})
}
