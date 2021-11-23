package usecases

import (
	accounts_usecase "github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
)

type AccountUseCase struct {
	storage accounts_usecase.Repository
}

func NewUseCase(storage accounts_usecase.Repository) AccountUseCase {
	return AccountUseCase{
		storage: storage,
	}
}
