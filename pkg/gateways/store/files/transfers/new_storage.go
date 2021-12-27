package transfers

import (
	"encoding/json"
	"log"
	"os"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type TransferRepository struct {
	storage   *os.File
	transfers map[string]entities.Transfer
}

func NewStorage(storage *os.File) TransferRepository {
	
	transferMap := make(map[string]entities.Transfer)
	readFile, err := os.ReadFile(storage.Name())
	if err != nil {
		log.Printf("error while reading file: %v", readFile)
		return TransferRepository{}
	}
	err = json.Unmarshal(readFile, &transferMap)
	if err != nil {
		log.Printf("error while deconding file: %v", err)
	}
	return TransferRepository{
		storage:   storage,
		transfers: transferMap,
	}
}
