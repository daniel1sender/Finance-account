package entities

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//Letramaiuscula
//descrição> Err,
var (
	ErrInvalidCPF      = errors.New("invalid informed cpf")
	ErrToGenerateHash  = errors.New("could not generate the hash")
	ErrInvalidName     = errors.New("empty informed name")
	ErrBalanceLessZero = errors.New("balance account is less than zero")
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

	if len(name) == 0 {
		return Account{}, ErrInvalidName
	}

	if len(cpf) != 11 {
		return Account{}, ErrInvalidCPF
	}

	hash, err := HashGenerator(secret)
	if err != nil {
		return Account{}, ErrToGenerateHash
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
		return "", fmt.Errorf("err to generate the hash %s", hash)
	}

	return string(hash), nil

}
