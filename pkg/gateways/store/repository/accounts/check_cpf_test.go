package accounts

import (
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func TestCheckCPF(t *testing.T) {

	t.Run("should return a null error when cpf account don't exist", func(t *testing.T) {
		storage := NewStorage()
		account := entities.Account{ID: "123", Balance: 10}
		err := storage.CheckCPF(account.CPF)
		if err != nil {
			t.Errorf("expected null error but got %v", err)
		}
	})

	t.Run("should return an error when cpf account exist", func(t *testing.T) {
		storage := NewStorage()
		account := entities.Account{ID: "123", Balance: 10}
		storage.users[account.ID] = account
		err := storage.CheckCPF(account.CPF)
		if err == nil {
			t.Errorf("expected null error but got %v", err)
		}
	})

}
