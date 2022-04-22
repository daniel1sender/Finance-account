package entities

import (
	"errors"
	"fmt"
	"time"

	"github.com/daniel1sender/Desafio-API/pkg/domain/verify"
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

	err := verify.IsValidSecret(secret)
	if err != nil {
		return Account{}, verify.ErrEmptySecret
	}

	err = verify.IsValidCPF(cpf)
	if err != nil {
		return Account{}, verify.ErrInvalidCPF
	}

	hash, err := HashGenerator(secret)
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

//essa função será resposável por criar o hash a partir do secret/password passado.
func HashGenerator(secret string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(secret), 4)

	if err != nil {
		return "", err
	}

	return string(hash), nil

}
