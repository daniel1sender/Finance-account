package transfers

import "github.com/daniel1sender/Desafio-API/pkg/domain/transfers"

type Handler struct {
	useCase transfers.UseCase
}

func NewHandler(useCase transfers.UseCase) Handler {
	return Handler{
		useCase: useCase,
	}
}
