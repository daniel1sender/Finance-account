package main

import (
	"fmt"

	"exemplo.com/domain"
)

func main() {
 	accounts := []domain.Account{
		{1, "daniel", "123", 10},
		{2, "erika", "12364", 0},
		{3, "daniel", "12345678910", 10},
		{4, "joão", "12345678910", 0},
	}

	for _, account := range accounts {
		accountMap, err := domain.CreateAccount(account)
		if err != nil {
			fmt.Printf("Deu ruim: %v\n", err)
			continue
		}
		// para saber mais sobre fmt: https://golang.org/pkg/fmt/
		fmt.Printf("Mapa de Contas: %+v\n", accountMap)
	} 

	listAccounts := domain.GetAccounts()

	fmt.Println(len(listAccounts))
	//fmt.Println(domain.GetBalanceById("3"))

	/* ---------------------------------------------------------------------------- */
	fmt.Println("make transfer:")

	transfer := domain.Transfer{"1", 3, 4, 20}
	
	fmt.Println(domain.MakeTransfer(transfer))
	
}
