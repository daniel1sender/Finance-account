package entities

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//passar a função hasheamento para a aprte de entidades
//time.Now().UTC()
type Account struct {
	Id        int
	Name      string
	Cpf       string
	Secret    string
	Balance   float64
	CreatedAt time.Time
}

func NewAccount(id int, name, cpf, secret string, balance float64) (Account, error) {

	if len(cpf) != 11 {
		return Account{}, fmt.Errorf("CPF %s is not correct", cpf)
	}

	hash, err := HashGenerator(secret)
	if err != nil {
		return Account{}, fmt.Errorf("err to generate the hash %s", hash)
	}

	return Account{
		Id:      id,
		Name:    name,
		Cpf:     cpf,
		Secret:  hash,
		Balance: balance,
		CreatedAt: time.Now().UTC(),
	}, nil
}



//essa função será resposável por criar o hash a partir do secret/password passado.
func HashGenerator(secret string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(secret), 4)

	if err != nil {
		return " ", fmt.Errorf("err to generate the hash %s", hash)
	}

	return string(hash), nil

}
