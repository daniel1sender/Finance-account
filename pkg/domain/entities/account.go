package entities

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCPF      = errors.New("cpf informed is invalid")
	ErrToGenerateHash  = errors.New("failed to process secret")
	ErrInvalidName     = errors.New("name informed is empty")
	ErrBalanceLessZero = errors.New("balance of the account created cannot be less than zero")
	ErrBlankSecret     = errors.New("secret informed is blanc")
)

type Account struct {
	ID        string
	Name      string
	CPF       string
	Secret    string
	Balance   int
	CreatedAt time.Time
}

func NewAccount(name, cpf, secret string, balance int) (Account, error) {

	if name == "" {
		return Account{}, ErrInvalidName
	}

	if len(cpf) != 11 {
		return Account{}, ErrInvalidCPF
	}

	if secret == "" {
		return Account{}, ErrBlankSecret
	}

	hash, err := HashGenerator(secret)
	if err != nil {
		return Account{}, fmt.Errorf("%w: %s", ErrToGenerateHash, err)
	}

	if balance < 0 {
		return Account{}, ErrBalanceLessZero
	}

	id := uuid.NewString()

	return Account{
		ID:        id,
		Name:      name,
		CPF:       cpf,
		Secret:    hash,
		Balance:   balance,
		CreatedAt: time.Now().UTC(),
	}, nil
}

//essa função será resposável por criar o hash a partir do secret/password passado.
func HashGenerator(secret string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(secret), 4)

	if err != nil {
		return "", err
	}

	return string(hash), nil

}
