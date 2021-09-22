package http

import "github.com/daniel1sender/Desafio-API/pkg/domain/accounts"

const (
	ContentType = "application/json"
)

type Handler struct {
	useCase accounts.AccountUseCase
}

func NewHandler(useCase accounts.AccountUseCase) Handler {
	return Handler{
		useCase: useCase,
	}
}
