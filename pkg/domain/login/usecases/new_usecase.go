package usecases

import "github.com/daniel1sender/Desafio-API/pkg/domain/login"

type LoginUseCase struct {
	AccountStorage  login.AccountRepository
	LoginRepository login.LoginRepository
	tokenSecret     string
}

func NewUseCase(accountStorage login.AccountRepository, loginRepository login.LoginRepository, tokenSecret string) LoginUseCase {
	return LoginUseCase{
		AccountStorage:  accountStorage,
		LoginRepository: loginRepository,
		tokenSecret:     tokenSecret,
	}
}
