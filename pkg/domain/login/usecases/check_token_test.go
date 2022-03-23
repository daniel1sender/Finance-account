package usecases

import (
	"context"
	"testing"
	"time"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/login"
	"github.com/daniel1sender/Desafio-API/pkg/tests"
	"github.com/stretchr/testify/assert"
)

func TestLoginUseCase_CheckToken(t *testing.T) {
	ctx := context.Background()
	accountRepository := accounts.NewStorage(Db)
	loginRepository := login.NewStorage(Db)
	tokenSecret := "123"
	useCase := LoginUseCase{accountRepository, loginRepository, tokenSecret}
	assert := assert.New(t)
	name := "Jonh Doe"
	cpf := "01481623559"
	accountSecret := "123"
	balance := 10
	account, _ := entities.NewAccount(name, cpf, accountSecret, balance)
	duration := "1m"
	accountRepository.Upsert(ctx, account)
	tokenString, _ := useCase.Auth(ctx, account.CPF, accountSecret, duration)
	expTime, _ := time.ParseDuration(duration)
	claims := entities.NewClaim(account.ID, expTime)

	t.Run("should return a null error when token exists in the database", func(t *testing.T) {
		loginRepository.Insert(ctx, claims, tokenString)
		err := useCase.CheckToken(ctx, tokenString)
		assert.Nil(err)
	})

	t.Run("should return an error when token does not exist in the database", func(t *testing.T) {
		tests.DeleteAllTokens(Db)
		err := useCase.CheckToken(ctx, tokenString)
		assert.NotNil(err)
	})
}
