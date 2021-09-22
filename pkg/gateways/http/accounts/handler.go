package accounts

import "github.com/daniel1sender/Desafio-API/pkg/domain/accounts"

type Error struct {
	Reason string `json:"reason"`
}

type Handler struct {
	useCase accounts.AccountUseCase
}

func NewHandler(useCase accounts.AccountUseCase) Handler {
	return Handler{
		useCase: useCase,
	}
}
