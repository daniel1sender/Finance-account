package main

import (
	"fmt"

	"exemplo.com/domain"
)

func main() {
	/* accounts := []domain.Account{
		domain.Account{Id: "1", Name: "daniel", Cpf: "123"},
		domain.Account{Id: "2", Name: "erika", Cpf: "12364"},
		domain.Account{Id: "3", Name: "daniel", Cpf: "12345678910"},
		domain.Account{Id: "4", Name: "joão", Cpf: "12345678910"},
	}

	for _, account := range accounts {
		accountMap, err := domain.CreateAccount(account)
		if err != nil {
			fmt.Printf("Deu ruim: %v\n", err)
			continue
		}
		// para saber mais sobre fmt: https://golang.org/pkg/fmt/
		fmt.Printf("Mapa de Contas: %+v\n", accountMap)
	} */

	mapAccounts := domain.GetAccounts()

	fmt.Println(mapAccounts)
	//fmt.Println(domain.GetBalanceById("3"))

}
