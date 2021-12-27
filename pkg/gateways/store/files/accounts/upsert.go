package accounts

import (
	"encoding/json"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (ar accountRepository) Upsert(account entities.Account) error {
	ar.users[account.ID] = entities.Account{ID: account.ID, Name: account.Name, CPF: account.CPF, Balance: account.Balance, CreatedAt: account.CreatedAt}
	keepAccount, err := json.MarshalIndent(ar.users, "", " ")
	if err != nil {
		return fmt.Errorf("error decoding account %v", err)
	}
	_, err = ar.storage.Write(keepAccount)
	if err != nil {	
		return fmt.Errorf("error while writing in file: %v", err)
	}
	return nil
}
