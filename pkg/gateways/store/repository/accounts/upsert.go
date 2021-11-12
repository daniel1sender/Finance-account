package accounts

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (ar accountRepository) Upsert(account entities.Account) {
	ar.users[account.ID] = entities.Account{ID: account.ID, Name: account.Name, CPF: account.CPF, Balance: account.Balance, CreatedAt: account.CreatedAt}
	keepAccount, err := json.MarshalIndent(ar.users, "", " ")
	if err != nil {
		log.Fatal("error decoding account")
	}
	_ = ioutil.WriteFile("Account_Repository.json", keepAccount, 0644)
}
