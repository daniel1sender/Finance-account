package usecases

import (
	"context"
	"errors"
	"testing"
	"time"

	accounts_usecases "github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/login"
	"github.com/daniel1sender/Desafio-API/pkg/tests"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestLoginUseCase_Login(t *testing.T) {
	ctx := context.Background()
	accountRepository := accounts.NewRepository(Db)
	loginRepository := login.NewRepository(Db)
	tokenSecret := "123"
	duration := "1m"
	useCase, _ := NewUseCase(accountRepository, loginRepository, tokenSecret, duration)
	assert := assert.New(t)

	t.Run("should return a signed token", func(t *testing.T) {
		name := "Jonh Doe"
		cpf := "01481623559"
		accountSecret := "123"
		balance := 10
		account, _ := entities.NewAccount(name, cpf, accountSecret, balance)
		accountRepository.Upsert(ctx, account)

		tokenString, err := useCase.Login(ctx, account.CPF, accountSecret)
		assert.Nil(err)
		assert.NotEmpty(tokenString)
		validateToken(t, tokenString, account.ID, tokenSecret)
		tests.DeleteAllAccounts(Db)
	})

	t.Run("should return an empty token and an error when account is not found", func(t *testing.T) {
		name := "Jonh Doe"
		cpf := "01481623550"
		accountSecret := "123"
		balance := 10
		account, _ := entities.NewAccount(name, cpf, accountSecret, balance)
		
		tokenString, err := useCase.Login(ctx, account.CPF, accountSecret)
		assert.True(errors.Is(err, accounts_usecases.ErrAccountNotFound))
		assert.Empty(tokenString)
		tests.DeleteAllAccounts(Db)
	})
}

func validateToken(t *testing.T, tokenString string, accountID string, tokenSecret string) {
	t.Helper()
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		t.Fatalf("expected no error but got '%v'", err)
	}
	claims := token.Claims.(*jwt.RegisteredClaims)
	assert.Equal(t, accountID, claims.Subject)
	assert.NotEmpty(t, claims.ID)
	if !claims.VerifyExpiresAt(time.Now(), true) {
		t.Error("expected non-zero 'expires at' time")
	}
	if !claims.VerifyIssuedAt(time.Now(), true) {
		t.Error("expected non-zero 'issued at' time")
	}
}
