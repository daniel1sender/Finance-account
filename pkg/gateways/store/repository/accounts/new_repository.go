package accounts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type accountRepository struct {
	storage *os.File
	users   map[string]entities.Account
}

func NewStorage() accountRepository {
	openFile, err := os.OpenFile("Account_Repository.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	accountMap := make(map[string]entities.Account)
	readFile, err := ioutil.ReadAll(openFile)
	if err != nil {
		return accountRepository{}
	}
	err = json.Unmarshal(readFile, &accountMap)
	if err != nil {
		fmt.Println(err)
	}
	return accountRepository{
		storage: openFile,
		users:   accountMap,
	}
}
