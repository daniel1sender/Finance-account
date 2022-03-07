package loginUseCases

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/daniel1sender/Desafio-API/pkg/config"
	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/accounts"
	"github.com/daniel1sender/Desafio-API/pkg/gateways/store/postgres/login"
	"github.com/daniel1sender/Desafio-API/pkg/tests"
)

func TestLoginUseCase_CheckToken(t *testing.T) {
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

	t.Run("should return a null error when token exists in the database", func(t *testing.T) {
		err := useCase.LoginStorage.Insert(ctx, token, config.TokenSecret)
		if err != nil {
			t.Errorf("expected not error but got '%s'", err.Error())
		}
		err = useCase.CheckToken(ctx, token)
		if err != nil {
			t.Errorf("expected not error but got '%s'", err.Error())
		}
	})

	t.Run("should return an error when token is not found", func(t *testing.T) {
		err := useCase.CheckToken(ctx, "token")
		if !errors.Is(err, login.ErrTokenNotFound) {
			t.Errorf("expected %v but got %v", login.ErrTokenNotFound, err)
		}
	})
	tests.DeleteAllTokens(Db)
}
