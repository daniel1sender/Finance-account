package entities

import (
	"errors"
	"testing"
)

func TestNewAccount(t *testing.T) {

	t.Run("Should successfully return an account", func(t *testing.T) {

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10
		account, err := NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("Expected nil but got %s", err)
		}

		if account.Name != name {
			t.Errorf("Expected %s but got %s", name, account.Name)
		}

		if account.CPF != cpf {
			t.Errorf("Expected %s but got %s", cpf, account.CPF)
		}

		if account.Balance != balance {
			t.Errorf("Expected %d but got %d", balance, account.Balance)
		}

		if account.CreatedAt.IsZero() {
			t.Errorf("Expected time different from zero time")
		}

		if account.Secret == secret {
			t.Error("Expected encrypted secret")
		}

	})

	t.Run("Should return a empty account and a error message when name is empty", func(t *testing.T) {

		name := ""
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := NewAccount(name, cpf, secret, balance)

		if account != (Account{}) {
			t.Errorf("Expected %+v but got %+v", Account{}, account)
		}

		if !errors.Is(err, ErrInvalidName) {
			t.Errorf("Expected error %s but got %s", ErrInvalidName, err)
		}

	})

	t.Run("Should return a empty account and a error message when cpf don't have 11 digits", func(t *testing.T) {

		name := "John Doe"
		cpf := "1111111030"
		secret := "123"
		balance := 10

		account, err := NewAccount(name, cpf, secret, balance)

		if account != (Account{}) {
			t.Errorf("Expected %v but got %v", account, Account{})
		}

		if !errors.Is(err, ErrInvalidCPF) {
			t.Errorf("Expected error %s but got %s", ErrInvalidCPF, err)
		}

	})

	t.Run("Should return a empty account and a error message when balance is less than zero", func(t *testing.T) {

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := -1

		account, err := NewAccount(name, cpf, secret, balance)

		if account != (Account{}) {
			t.Errorf("Expected %v but got %v", account, Account{})
		}

		if !errors.Is(err, ErrBalanceLessZero) {
			t.Errorf("Expected error %s but got %s", ErrBalanceLessZero, err)
		}

	})

}
