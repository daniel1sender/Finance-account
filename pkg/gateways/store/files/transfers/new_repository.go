package transfers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type TransferRepository struct {
	storage *os.File
	users   map[string]entities.Transfer
}

func NewRepository() TransferRepository {
	openFile, err := os.OpenFile("Transfer_Respository.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	transferMap := make(map[string]entities.Transfer)
	readFile, err := ioutil.ReadAll(openFile)
	if err != nil {
		return TransferRepository{}
	}
	err = json.Unmarshal(readFile, &transferMap)
	if err != nil {
		fmt.Println(err)
	}
	return TransferRepository{
		storage: openFile,
		users:   transferMap,
	}
}
