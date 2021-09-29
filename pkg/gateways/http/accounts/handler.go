package accounts

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
)

type Error struct {
	Reason string `json:"reason"`
}

type Handler struct {
	useCase accounts.UseCase
}

func NewHandler(useCase accounts.UseCase) Handler {
	return Handler{
		useCase: useCase,
	}
}
