package entities

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//passar a função hasheamento para a aprte de entidades
//time.Now().UTC()
//mudar o id

type Account struct {
	Id        string
	Name      string
	Cpf       string
	Secret    string
	Balance   int
	CreatedAt time.Time
}

func NewAccount(name, cpf, secret string, balance int) (Account, error) {

	if len(cpf) != 11 {
		return Account{}, fmt.Errorf("CPF %s is not correct", cpf)
	}

	hash, err := HashGenerator(secret)
	if err != nil {
		return Account{}, fmt.Errorf("err to generate the hash %s", hash)
	}

	id := uuid.NewString()

	return Account{
		Id:        id,
		Name:      name,
		Cpf:       cpf,
		Secret:    hash,
		Balance:   balance,
		CreatedAt: time.Now().UTC(),
	}, nil
}

//essa função será resposável por criar o hash a partir do secret/password passado.
//comparar se o hash gerado é diferente do secret

func HashGenerator(secret string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(secret), 4)

	if err != nil {
		return " ", fmt.Errorf("err to generate the hash %s", hash)
	}

	//passar para os testes
	/* 	err = bcrypt.CompareHashAndPassword(hash, []byte(secret))

	   	if err != nil {
	   		return " ", fmt.Errorf("hash: %s equal secret: %s", hash, secret)
	   	} */

	return string(hash), nil

}
