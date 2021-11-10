package accounts

import (
	"log"
	"os"
	"time"
)

type AccountRepository struct {
	storage *os.File
}

type AccountResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CPF       string    `json:"cpf"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func NewStorage() AccountRepository {
	file, err := os.OpenFile("repository", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return AccountRepository{
		storage: file,
	}
}
