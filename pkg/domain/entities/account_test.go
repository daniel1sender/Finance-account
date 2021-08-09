package entities

import (
	"testing"
)

//fazer os casos positivos e negativos da newaccount
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

		if account.Cpf != cpf {
			t.Errorf("Expected %s but got %s", cpf, account.Cpf)
		}

		if account.Balance != balance {
			t.Errorf("Expected %d but got %d", balance, account.Balance)
		}

		if account.CreatedAt.IsZero() == true {
			t.Errorf("Expected time different from zero time")
		}

		if account.Secret == secret {
			t.Error("Expected incripted secret")
		}

	})

	t.Run("Should return a empty account and a error message when name is empty", func(t *testing.T) {

		name := ""
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := NewAccount(name, cpf, secret, balance)
		if err == nil {
			t.Errorf("Expected nil but got %s", err)
		}

		if account.Name != name {
			t.Errorf("Expected %s but got %s", name, account.Name)
		}

		if account.Cpf == cpf {
			t.Errorf("Expected %s but got %s", cpf, account.Cpf)
		}

		if account.Balance == balance {
			t.Errorf("Expected %d but got %d", balance, account.Balance)
		}

		if account.CreatedAt.IsZero() != true {
			t.Errorf("Expected time different from zero time")
		}

		if account.Secret == secret {
			t.Error("Expected incripted secret")
		}

	})

	t.Run("Should return a empty account and a error message when cpf don't have 11 digits", func(t *testing.T) {

		name := "John Doe"
		cpf := "1111111030"
		secret := "123"
		balance := 10

		account, err := NewAccount(name, cpf, secret, balance)
		if err == nil {
			t.Errorf("Expected nil but got %s", err)
		}

		if account.Name == name {
			t.Errorf("Expected %s but got %s", name, account.Name)
		}

		if account.Cpf == cpf {
			t.Errorf("Expected %s but got %s", cpf, account.Cpf)
		}

		if account.Balance == balance {
			t.Errorf("Expected %d but got %d", balance, account.Balance)
		}

		if account.CreatedAt.IsZero() != true {
			t.Errorf("Expected time different from zero time")
		}

		if account.Secret == secret {
			t.Error("Expected incripted secret")
		}

	})

	t.Run("Should return a empty account and a error message when occour error to generate the hash", func(t *testing.T) {

		name := "John Doe"
		cpf := "1111111030"
		secret := "123"
		balance := 10

		account, err := NewAccount(name, cpf, secret, balance)
		if err == nil {
			t.Errorf("Expected nil but got %s", err)
		}

		if account.Name == name {
			t.Errorf("Expected %s but got %s", name, account.Name)
		}

		if account.Cpf == cpf {
			t.Errorf("Expected %s but got %s", cpf, account.Cpf)
		}

		if account.Balance == balance {
			t.Errorf("Expected %d but got %d", balance, account.Balance)
		}

		if account.CreatedAt.IsZero() != true {
			t.Errorf("Expected time different from zero time")
		}

		if account.Secret == secret {
			t.Error("Expected incripted secret")
		}

	})

	t.Run("Should return a empty account and a error message when balance is less or equal zero", func(t *testing.T) {

		name := "John Doe"
		cpf := "1111111030"
		secret := "123"
		balance := -1

		account, err := NewAccount(name, cpf, secret, balance)
		if err == nil {
			t.Errorf("Expected nil but got %s", err)
		}

		if account.Name == name {
			t.Errorf("Expected %s but got %s", name, account.Name)
		}

		if account.Cpf == cpf {
			t.Errorf("Expected %s but got %s", cpf, account.Cpf)
		}

		if account.Balance == balance {
			t.Errorf("Expected %d but got %d", balance, account.Balance)
		}

		if account.CreatedAt.IsZero() != true {
			t.Errorf("Expected time different from zero time")
		}

		if account.Secret == secret {
			t.Error("Expected incripted secret")
		}

	})

}
