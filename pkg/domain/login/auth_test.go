package login

import (
	"context"
	"testing"
	"time"

	accounts_usecases "github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/tests"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestLoginUseCase_Auth(t *testing.T) {
	ctx := context.Background()
	accountRepository := accounts.NewStorage(Db)
	tokenSecret := "123"
	useCase := LoginUseCase{accountRepository, tokenSecret}
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
		assert.NoError(err)
		assert.NotEmpty(tokenString, "got empty token")
		validateToken(t, tokenString, account.ID, tokenSecret)
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
		assert.NotEqualf(err, accounts_usecases.ErrAccountNotFound, "expected '%v' but got '%v'", accounts_usecases.ErrAccountNotFound, err)
		assert.Empty(tokenString, "got a non-empty token")
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
	assert.Equalf(t, claims.Subject, accountID, "expected '%s' but got '%s'", accountID, claims.Subject)
	assert.NotEqual(t, claims.ID, "", "expected not empty id")

	if !claims.VerifyExpiresAt(time.Now(), true) {
		t.Error("expected non-zero 'expires at' time")
	}
	if !claims.VerifyIssuedAt(time.Now(), true) {
		t.Error("expected non-zero 'issued at' time")
	}
}
