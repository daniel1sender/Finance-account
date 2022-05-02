package entities

import (
	"errors"
	"fmt"
	"time"

	"github.com/daniel1sender/Desafio-API/pkg/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrToGenerateHash  = errors.New("failed to process secret")
	ErrInvalidName     = errors.New("name informed is empty")
	ErrNegativeBalance = errors.New("balance of the account created cannot be less than zero")
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

	err := domain.ValidateSecret(secret)
	if err != nil {
		return Account{}, fmt.Errorf("error while validating secret: %w", err)
	}

	err = domain.ValidateCPF(cpf)
	if err != nil {
		return Account{}, fmt.Errorf("error while validating cpf: %w", err)
	}

	hash, err := GenerateHash(secret)
	if err != nil {
		return Account{}, fmt.Errorf("%w: %s", ErrToGenerateHash, err)
	}

	if balance < 0 {
		return Account{}, ErrNegativeBalance
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

func GenerateHash(secret string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(secret), 4)

	if err != nil {
		return "", err
	}

	return string(hash), nil

}
