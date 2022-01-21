package http

import "github.com/sirupsen/logrus"

const (
	JSONContentType = "application/json"
	DateLayout      = "2006-01-02T15:04:05Z"
	KeyHeader       = "request-id"
)

type Error struct {
	Reason string `json:"reason"`
}

func NewLogger() *logrus.Entry {
	log := logrus.NewEntry(logrus.New())
	log.Logger.SetFormatter(&logrus.JSONFormatter{})
	return log
}
