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

func NewStorage() TransferRepository {
	fileName := "Transfer_Respository.json"
	openFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error to open file: %v", err)
	}
	trasnferMap := make(map[string]entities.Transfer)
	readFile, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("error while reading file: %v", readFile)
		return TransferRepository{}
	}
	err = json.Unmarshal(readFile, &trasnferMap)
	if err != nil {
		log.Fatalf("error while deconding file: %v", err)
	}
	return TransferRepository{
		storage:   openFile,
		transfers: trasnferMap,
	}
}
