package accounts

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type accountRepository struct {
	users map[string]entities.Account
}

func NewStorage() accountRepository {
	accountMap := make(map[string]entities.Account)
	if _, err := os.Stat("Account_Repository.json"); err != nil {
		_, err := os.OpenFile("Account_Repository.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error to open file: %v", err)
		}
	}
	readFile, err := os.ReadFile("Account_Repository.json")
	if err != nil {
		return accountRepository{}
	}
	err = json.Unmarshal(readFile, &accountMap)
	if err != nil {
		fmt.Println(err)
	}
	return accountRepository{
		users: accountMap,
	}
}
