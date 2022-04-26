package transfers

import "github.com/daniel1sender/Desafio-API/pkg/domain/entities"

type TransferRepository struct {
	storage map[string]entities.Transfer
}

func NewRepository() TransferRepository {
	sto := make(map[string]entities.Transfer)
	return TransferRepository{
		storage: sto,
	}
}
