package accounts

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
)

type Handler struct {
	useCase accounts.UseCase
}

func NewHandler(useCase accounts.UseCase) Handler {
	return Handler{
		useCase: useCase,
	}
}
