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

func TestLoginUsecase_ValidateToken(t *testing.T) {
	ctx := context.Background()
	accountRepository := accounts.NewRepository(Db)
	loginRepository := login.NewRepository(Db)
	tokenSecret := "123"
	duration := "1m"
	useCase, _ := NewUseCase(accountRepository, loginRepository, tokenSecret, duration)
	name := "Jonh Doe"
	cpf := "01481623559"
	accountSecret := "123"
	balance := 10
	account, _ := entities.NewAccount(name, cpf, accountSecret, balance)
	newClaims := entities.NewClaim(account.ID, useCase.expTime)
	tokenString, _ := GenerateJWT(newClaims, tokenSecret)

	t.Run("should return a claim successfully", func(t *testing.T) {
		loginRepository.Insert(ctx, newClaims, tokenString)
		claim, err := useCase.ValidateToken(ctx, tokenString)
		assert.NotEmpty(t, claim.TokenID)
		assert.NotEmpty(t, claim.Sub)
		assert.NotEmpty(t, claim.ExpTime)
		assert.NotEmpty(t, claim.CreatedTime)
		assert.Equal(t, newClaims.TokenID, claim.TokenID)
		assert.Equal(t, newClaims.Sub, claim.Sub)
		assert.NoError(t, err)
	})

	t.Run("should return an empty claim and an error when token is not found", func(t *testing.T) {
		tests.DeleteAllTokens(Db)
		claim, err := useCase.ValidateToken(ctx, tokenString)
		assert.Error(t, err)
		assert.Empty(t, claim)
	})

	t.Run("should return an empty claim and an error when the token was expired", func(t *testing.T) {
		duration := "1ms"
		expTime, _ := time.ParseDuration(duration)
		newClaim := entities.NewClaim(account.ID, expTime)
		tokenString, _ := GenerateJWT(newClaims, tokenSecret)
		loginRepository.Insert(ctx, newClaim, tokenString)
		claims, err := useCase.ValidateToken(ctx, tokenString)
		assert.Error(t, err)
		assert.Empty(t, claims)
	})

	t.Run("should return an empty claim and an error because token signature method is invalid", func(t *testing.T) {
		tokenString := "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.tyh-VfuzIxCyGYDlkBA7DfyjrqmSHu6pQ2hoZuFqUSLPNY2N0mpHb3nk5K17HWP_3cYHBw7AhHale5wky6-sVA"
		claims, err := useCase.ValidateToken(ctx, tokenString)
		assert.Error(t, err)
		assert.Empty(t, claims)
	})
}
