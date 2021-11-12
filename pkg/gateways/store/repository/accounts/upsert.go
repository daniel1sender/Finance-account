package accounts

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (ar AccountRepository) Upsert(account entities.Account) {
	users := ar.Users
	users[account.ID] = entities.Account{Name: account.Name, CPF: account.CPF, Balance: account.Balance, CreatedAt: account.CreatedAt}
	keepAccount, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		log.Fatal("error decoding account")
	}
	_ = ioutil.WriteFile("repository.json", keepAccount, 0644)
}
