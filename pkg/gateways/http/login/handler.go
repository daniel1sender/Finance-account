package login

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/login"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	UseCase login.UseCase
	logger  *logrus.Entry
}

func NewHandler(useCase login.UseCase, logger *logrus.Entry) Handler {
	return Handler{
		UseCase: useCase,
		logger:  logger,
	}
}
