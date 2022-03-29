package usecases

import (
	"fmt"
	"time"

	"github.com/daniel1sender/Desafio-API/pkg/domain/login"
)

type LoginUseCase struct {
	AccountStorage  login.AccountRepository
	LoginRepository login.Repository
	tokenSecret     string
	expTime         time.Duration
}

func NewUseCase(accountStorage login.AccountRepository, loginRepository login.Repository, tokenSecret, duration string) (LoginUseCase, error) {
	expTime, err := time.ParseDuration(duration)
	if err != nil {
		return LoginUseCase{}, fmt.Errorf("error while parsing duration time: %w", err)
	}
	return LoginUseCase{
		AccountStorage:  accountStorage,
		LoginRepository: loginRepository,
		tokenSecret:     tokenSecret,
		expTime:         expTime,
	}, nil
}
