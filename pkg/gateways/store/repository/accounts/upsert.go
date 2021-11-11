package accounts

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (ar AccountRepository) Upsert(account entities.Account) {
	users := ar.Users
	users[account.ID] = AccountResponse{account.Name, account.CPF, account.Balance, account.CreatedAt}

	keepAccount, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		log.Fatal("error decoding account")
	}
	//ar.storage.WriteString(string(keepAccount) + "\n")
	_ = ioutil.WriteFile("repository.json", keepAccount, 0644)
}
