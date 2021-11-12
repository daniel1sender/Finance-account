package accounts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type AccountRepository struct {
	storage *os.File
	Users   map[string]entities.Account
}

func NewStorage() AccountRepository {
	openFile, err := os.OpenFile("repository.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	accountMap := make(map[string]entities.Account)

	readFile, err := ioutil.ReadAll(openFile)
	if err != nil {
		return AccountRepository{}
	}

	err = json.Unmarshal(readFile, &accountMap)
	if err != nil {
		fmt.Println(err)
	}

	return AccountRepository{
		storage: openFile,
		Users:   accountMap,
	}
}
