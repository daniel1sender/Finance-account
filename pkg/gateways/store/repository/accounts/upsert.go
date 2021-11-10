package accounts

import (
	"encoding/json"
	"log"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (ar AccountRepository) Upsert(account entities.Account) {
	accountResponse := AccountResponse{account.ID, account.Name, account.CPF, account.Balance, account.CreatedAt}
	keepAccount, err := json.Marshal(accountResponse)
	if err != nil {
		log.Fatal("error decoding account")
	}
	ar.storage.WriteString(string(keepAccount) + "\n")
}
