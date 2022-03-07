package login

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/login"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	useCase login.UseCase
	logger  *logrus.Entry
}

func NewHandler(useCase login.LoginUseCase, logger *logrus.Entry) Handler {
	return Handler{
		useCase: useCase,
		logger:  logger,
	}
}
