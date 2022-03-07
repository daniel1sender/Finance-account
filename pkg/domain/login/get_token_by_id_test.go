package login

import (
	"context"
	"testing"
	"time"

	"github.com/daniel1sender/Desafio-API/pkg/config"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/login"
	"github.com/daniel1sender/Desafio-API/pkg/tests"
)

func TestLoginUseCase_GetTokenByID(t *testing.T) {
	ctx := context.Background()
	accountRepository := accounts.NewStorage(Db)
	loginRepository := login.NewStorage(Db)
	tokenSecret := "123"
	duration, _ := time.ParseDuration("2m")
	useCase := LoginUseCase{loginRepository, accountRepository, tokenSecret}
	config, _ := config.GetConfig()
	account, _ := entities.NewAccount("Jonh Doe", "12345678910", "123", 10)
	tokenJWT, _ := GenerateJWT(account.ID, config.TokenSecret, duration)
	token, _ := tokenJWT.SignedString([]byte(config.TokenSecret))

	t.Run("should return a token", func(t *testing.T) {
		err := useCase.LoginStorage.Insert(ctx, token, config.TokenSecret)
		if err != nil {
			t.Errorf("expected no error but got '%s'", err.Error())
		}
		token, err := useCase.GetTokenByID(ctx, account.ID)
		if err != nil {
			t.Errorf("expected no error but got '%s'", err.Error())
		}
		err = ValidateToken(token, account.ID, config.TokenSecret)
		if err != nil {
			t.Errorf("expected no error but got '%s'", err.Error())
		}

	})

	t.Run("expected an empty token and an error when the token does not exist in the database", func(t *testing.T) {
		tests.DeleteAllTokens(Db)
		token, err := useCase.GetTokenByID(ctx, account.ID)
		if err == nil {
			t.Error("expected an error but got nil")
		}
		err = ValidateToken(token, account.ID, config.TokenSecret)
		if err == nil {
			t.Error("expected an error but got nil")
		}
	})
}
