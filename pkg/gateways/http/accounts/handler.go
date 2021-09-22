package accounts

import "github.com/daniel1sender/Desafio-API/pkg/domain/accounts"

type Handler struct {
	useCase accounts.AccountUseCase
}

func NewHandler(useCase accounts.AccountUseCase) Handler {
	return Handler{
		useCase: useCase,
	}
}
