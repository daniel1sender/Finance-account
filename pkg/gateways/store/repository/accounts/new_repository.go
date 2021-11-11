package accounts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type AccountRepository struct {
	storage *os.File
	Users   map[string]AccountResponse
}

type AccountResponse struct {
	//	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CPF       string    `json:"cpf"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type Account struct {
	Users map[string]AccountResponse
}

func NewStorage() AccountRepository {
	openFile, err := os.OpenFile("repository.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	accountMap := make(map[string]AccountResponse)

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
