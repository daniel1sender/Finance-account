package usecases

import (
	"context"
	"testing"
	"time"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/login"
	"github.com/daniel1sender/Desafio-API/pkg/tests"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestLoginUsecase_ValidateToken(t *testing.T) {
	ctx := context.Background()
	accountRepository := accounts.NewStorage(Db)
	loginRepository := login.NewStorage(Db)
	tokenSecret := "123"
	duration := "1m"
	useCase := LoginUseCase{accountRepository, loginRepository, tokenSecret, duration}
	name := "Jonh Doe"
	cpf := "01481623559"
	accountSecret := "123"
	balance := 10
	account, _ := entities.NewAccount(name, cpf, accountSecret, balance)
	expTime, _ := time.ParseDuration(useCase.expTime)
	claim := entities.NewClaim(account.ID, expTime)
	claims := jwt.RegisteredClaims{
		Subject:   account.ID,
		ExpiresAt: jwt.NewNumericDate(claim.ExpTime),
		IssuedAt:  jwt.NewNumericDate(claim.CreatedTime),
		ID:        claim.TokenID,
	}
	tokenString, _ := GenerateJWT(claims, tokenSecret)

	t.Run("should return a token succesfully", func(t *testing.T) {
		loginRepository.Insert(ctx, claim, tokenString)
		token, err := useCase.ValidateToken(ctx, tokenString)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("should return an empty token and an error when token is not found", func(t *testing.T) {
		tests.DeleteAllTokens(Db)
		token, err := useCase.ValidateToken(ctx, tokenString)
		assert.Error(t, err)
		assert.Empty(t, token)
	})

	t.Run("should return an empty token and an error when the token was expired", func(t *testing.T) {
		duration := "1ms"
		expTime, _ := time.ParseDuration(duration)
		claim := entities.NewClaim(account.ID, expTime)
		claims := jwt.RegisteredClaims{
			Subject:   account.ID,
			ExpiresAt: jwt.NewNumericDate(claim.ExpTime),
			IssuedAt:  jwt.NewNumericDate(claim.CreatedTime),
			ID:        claim.TokenID,
		}
		tokenString, _ := GenerateJWT(claims, tokenSecret)
		loginRepository.Insert(ctx, claim, tokenString)
		token, err := useCase.ValidateToken(ctx, tokenString)
		assert.Error(t, err)
		assert.Empty(t, token)
	})

	t.Run("should return an empty token and an error because token signature method is invalid", func(t *testing.T) {
		tokenString := "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.tyh-VfuzIxCyGYDlkBA7DfyjrqmSHu6pQ2hoZuFqUSLPNY2N0mpHb3nk5K17HWP_3cYHBw7AhHale5wky6-sVA"
		token, err := useCase.ValidateToken(ctx, tokenString)
		assert.Error(t, err)
		assert.Empty(t, token)
	})
}
