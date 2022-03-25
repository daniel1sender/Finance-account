package usecases

import "github.com/daniel1sender/Desafio-API/pkg/domain/login"

type LoginUseCase struct {
	AccountStorage  login.AccountRepository
	LoginRepository login.Repository
	tokenSecret     string
}

func NewUseCase(accountStorage login.AccountRepository, loginRepository login.Repository, tokenSecret string) LoginUseCase {
	return LoginUseCase{
		AccountStorage:  accountStorage,
		LoginRepository: loginRepository,
		tokenSecret:     tokenSecret,
	}
}
