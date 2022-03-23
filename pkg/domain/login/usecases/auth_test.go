package usecases

import (
	"context"
	"errors"
	"testing"

	accounts_usecases "github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/login"
	"github.com/daniel1sender/Desafio-API/pkg/tests"
	"github.com/stretchr/testify/assert"
)

func TestLoginUseCase_Auth(t *testing.T) {
	ctx := context.Background()
	accountRepository := accounts.NewStorage(Db)
	loginRepository := login.NewStorage(Db)
	tokenSecret := "123"
	useCase := LoginUseCase{accountRepository, loginRepository, tokenSecret}
	assert := assert.New(t)

	t.Run("should return a signed token", func(t *testing.T) {
		name := "Jonh Doe"
		cpf := "01481623559"
		accountSecret := "123"
		balance := 10
		account, _ := entities.NewAccount(name, cpf, accountSecret, balance)
		duration := "1m"
		accountRepository.Upsert(ctx, account)
		tokenString, err := useCase.Auth(ctx, account.CPF, accountSecret, duration)
		assert.Nil(err)
		assert.NotEmpty(tokenString)
		err = tests.ValidateToken(tokenString, account.ID, tokenSecret)
		assert.Nil(err)
		tests.DeleteAllAccounts(Db)
	})

	t.Run("should return an empty token and an error when account is not found", func(t *testing.T) {
		name := "Jonh Doe"
		cpf := "01481623550"
		accountSecret := "123"
		balance := 10
		account, _ := entities.NewAccount(name, cpf, accountSecret, balance)
		duration := "1m"
		tokenString, err := useCase.Auth(ctx, account.CPF, accountSecret, duration)
		assert.True(errors.Is(err, accounts_usecases.ErrAccountNotFound))
		assert.Empty(tokenString)
		tests.DeleteAllAccounts(Db)
	})
}
