package loginUseCases

import "github.com/daniel1sender/Desafio-API/pkg/domain/login"

type LoginUseCase struct {
	LoginStorage   login.Repository
	AccountStorage login.AccountRepository
	tokenSecret    string
}

func NewUseCase(loginStorage login.Repository, accountStorage login.AccountRepository, tokenSecret string) LoginUseCase {
	return LoginUseCase{
		LoginStorage:   loginStorage,
		AccountStorage: accountStorage,
		tokenSecret:    tokenSecret,
	}
}
