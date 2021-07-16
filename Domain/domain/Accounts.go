package domain

import "fmt"

type Account struct {
	Id      string
	Name    string
	Cpf     string
	Balance int
}

type AccountCreator struct {
}

var AccountNumber int = 0
var m = make(map[int]Account)

//Função para criar conta
func CreateAccount(a Account) (map[int]Account, error) {

	if len(a.Cpf) != 11 {
		return nil, fmt.Errorf("CPF %s is not correct", a.Cpf)
	}

	m[AccountNumber] = a
	AccountNumber++

	return m, fmt.Errorf("Cpf is %s correct", a.Cpf)
}

func GetAccounts() []Account {
	fmt.Printf("Accounts Created\n")
	var AccountsList []Account

	for _, value := range m {
		AccountsList = append(AccountsList, value)
	}
	return AccountsList
}

//erros possíveis em GetBalance, caso um Id inesistente seja passado
func GetBalanceById(id string) (int, error) {

	fmt.Println("Returned balance by ID")

	for key, value := range m {
		if value.Id == id {
			balance := m[key].Balance
			return balance, nil
		}
	}
	return 0, fmt.Errorf("No id %s found", id)
}
