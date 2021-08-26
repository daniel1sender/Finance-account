package entities

import (
	"errors"
	"testing"
)

func TestNewAccount(t *testing.T) {

	t.Run("should successfully return an account", func(t *testing.T) {

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10
		account, err := NewAccount(name, cpf, secret, balance)
		if err != nil {
			t.Errorf("expected no error but got '%s'", err)
		}

		if account.Name != name {
			t.Errorf("expected '%s' but got '%s'", name, account.Name)
		}

		if account.CPF != cpf {
			t.Errorf("expected '%s' but got '%s'", cpf, account.CPF)
		}

		if account.Balance != balance {
			t.Errorf("expected '%d' but got '%d'", balance, account.Balance)
		}

		if account.CreatedAt.IsZero() {
			t.Errorf("expected time different from zero time")
		}

		if account.Secret == secret {
			t.Error("expected encrypted secret")
		}

	})

	t.Run("should return a empty account and a error message when name is empty", func(t *testing.T) {

		name := ""
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := NewAccount(name, cpf, secret, balance)

		if account != (Account{}) {
			t.Errorf("expected '%+v' but got '%+v'", Account{}, account)
		}

		if !errors.Is(err, ErrInvalidName) {
			t.Errorf("expected error '%s' but got '%s'", ErrInvalidName, err)
		}

	})

	t.Run("should return a empty account and a error message when cpf don't have 11 digits", func(t *testing.T) {

		name := "John Doe"
		cpf := "1111111030"
		secret := "123"
		balance := 10

		account, err := NewAccount(name, cpf, secret, balance)

		if account != (Account{}) {
			t.Errorf("expected '%+v' but got '%+v'", account, Account{})
		}

		if !errors.Is(err, ErrInvalidCPF) {
			t.Errorf("expected error '%s' but got '%s'", ErrInvalidCPF, err)
		}

	})

	t.Run("should return a blank account and a error message when secret informed is empty", func(t *testing.T) {

		name := "John Doe"
		cpf := "11111111030"
		secret := ""
		balance := 10

		account, err := NewAccount(name, cpf, secret, balance)

		if account != (Account{}) {
			t.Errorf("expected '%+v' but got '%+v'", Account{}, account)
		}

		if !errors.Is(err, ErrBlankSecret) {
			t.Errorf("expected '%s' but got '%s'", ErrBlankSecret, err)
		}

	})

	t.Run("should return a empty account and a error message when balance is less than zero", func(t *testing.T) {

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := -1

		account, err := NewAccount(name, cpf, secret, balance)

		if account != (Account{}) {
			t.Errorf("expected '%+v' but got '%+v'", account, Account{})
		}

		if !errors.Is(err, ErrBalanceLessZero) {
			t.Errorf("expected error '%s' but got '%s'", ErrBalanceLessZero, err)
		}

	})

}
