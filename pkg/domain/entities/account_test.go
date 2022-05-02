package entities

import (
	"testing"

	"github.com/daniel1sender/Desafio-API/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {

	t.Run("should successfully return an account", func(t *testing.T) {

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := 10
		account, err := NewAccount(name, cpf, secret, balance)

		assert.Nil(t, err)
		assert.Equal(t, account.Name, name)
		assert.Equal(t, account.CPF, cpf)
		assert.Equal(t, account.Balance, balance)
		assert.NotEmpty(t, account.CreatedAt)
	})

	t.Run("should return an empty account and an error when name is empty", func(t *testing.T) {

		name := ""
		cpf := "11111111030"
		secret := "123"
		balance := 10

		account, err := NewAccount(name, cpf, secret, balance)

		assert.Empty(t, account)
		assert.Equal(t, err, ErrInvalidName)
	})

	t.Run("should return an empty account and an error when cpf doesn't have 11 digits", func(t *testing.T) {

		name := "John Doe"
		cpf := "1111111030"
		secret := "123"
		balance := 10

		account, err := NewAccount(name, cpf, secret, balance)

		assert.Empty(t, account)
		assert.Equal(t, err, domain.ErrInvalidCPF)
	})

	t.Run("should return an empty account and an error when secret informed is empty", func(t *testing.T) {

		name := "John Doe"
		cpf := "11111111030"
		secret := ""
		balance := 10

		account, err := NewAccount(name, cpf, secret, balance)

		assert.Empty(t, account)
		assert.Equal(t, err, domain.ErrEmptySecret)
	})

	t.Run("should return an empty account and an error when balance is less than zero", func(t *testing.T) {

		name := "John Doe"
		cpf := "11111111030"
		secret := "123"
		balance := -1

		account, err := NewAccount(name, cpf, secret, balance)

		assert.Empty(t, account)
		assert.Equal(t, err, ErrNegativeBalance)
	})

}
