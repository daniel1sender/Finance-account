package domain

import "fmt"

type Account struct {
	Id      int
	Name    string
	Cpf     string
	Balance float64
}

type AccountCreator struct {
}

var AccountNumber int = 0
var AccountsMap = make(map[int]Account)

//Função para criar conta
func CreateAccount(a Account) (map[int]Account, error) {

	if len(a.Cpf) != 11 {
		return nil, fmt.Errorf("CPF %s is not correct", a.Cpf)
	}

	a.Id = AccountNumber
	AccountsMap[AccountNumber] = a
	AccountNumber++

	return AccountsMap, nil
}

func GetAccounts() []Account {
	fmt.Printf("Accounts Created\n")
	var AccountsList []Account

	for _, value := range AccountsMap {
		AccountsList = append(AccountsList, value)
	}

	return AccountsList
}

//erros possíveis em GetBalance, caso um Id inesistente seja passado
func GetBalanceById(id int) (float64, error) {

	fmt.Println("Returned balance by ID")

	for key, value := range AccountsMap {
		if value.Id == id {
			balance := AccountsMap[key].Balance
			return balance, nil
		}
	}
	return 0, fmt.Errorf("no id %d found", id)
}

func AccountTransfer(t Transfer)(Account, Account, error) {
	//aqui preciso alterar o valor da conta e não do map

	fmt.Println("TRANSFER ACCOUNT")

	origin := t.Account_origin_id
	destination := t.Account_destinantion_id
	amount := t.Amount
fmt.Println("origin: ", origin, "destination: ", destination)

	balanceOrigin := AccountsMap[origin].Balance
	balanceDestination := AccountsMap[destination].Balance
	fmt.Println("BalanceOrigin: ", balanceOrigin, "BalanceDestination: ", balanceDestination)

	if (balanceOrigin - amount) < 0 {
		return AccountsMap[origin], AccountsMap[destination], fmt.Errorf("the origin account does not have enough balance")
	}

	balanceOrigin -= amount
	balanceDestination += amount

	accountOrigin := AccountsMap[origin]
	accountDestination := AccountsMap[destination]

	accountOrigin.Balance = balanceOrigin
	accountDestination.Balance = balanceDestination

	AccountsMap[origin] = accountOrigin
	AccountsMap[destination] = accountDestination

	return accountOrigin, accountDestination, nil
}  