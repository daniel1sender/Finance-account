package accounts

import (
	"encoding/json"
	"log"
	"os"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type accountRepository struct {
	storage *os.File
	users   map[string]entities.Account
}

func NewStorage(storage *os.File) accountRepository {
	accountMap := make(map[string]entities.Account)

	readFile, err := os.ReadFile(storage.Name())
	if err != nil {
		log.Printf("error while reading file: %v", err)
		return accountRepository{}
	}
	err = json.Unmarshal(readFile, &accountMap)
	if err != nil {
		log.Printf("error while decoding file: %v", err)
	}
	return accountRepository{
		users: accountMap,
	}
}
