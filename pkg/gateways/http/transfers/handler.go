package transfers

import "github.com/daniel1sender/Desafio-API/pkg/domain/transfers"

type Handler struct {
	useCase transfers.TransferUseCase
}

func NewHandler(useCase transfers.TransferUseCase) Handler {
	return Handler{
		useCase: useCase,
	}
}
