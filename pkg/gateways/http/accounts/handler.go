package accounts

import (
	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	useCase accounts.UseCase
	logger  *logrus.Entry
}

func NewHandler(useCase accounts.UseCase, logger *logrus.Entry) Handler {
	return Handler{
		useCase: useCase,
		logger:  logrus.NewEntry(logrus.New()),
	}
}
