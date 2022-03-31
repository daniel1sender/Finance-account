package usecases

import (
	"context"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/domain/login"
	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (l LoginUseCase) Login(ctx context.Context, cpf, accountSecret string) (string, error) {
	if len(accountSecret) == 0 {
		return "", fmt.Errorf("error while validating secret informed: %w", login.ErrEmptySecret)
	}
	if len(cpf) != 11 {
		return "", fmt.Errorf("error while validanting cpf informed: %w", login.ErrInvalidCPF)
	}
	account, err := l.AccountStorage.GetByCPF(ctx, cpf)
	if err != nil {
		return "", fmt.Errorf("error while getting account by cpf: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Secret), []byte(accountSecret))
	if err != nil {
		return "", fmt.Errorf("error while comparing secret informed with stored: %w", login.ErrInvalidSecret)
	}

	claim := entities.NewClaim(account.ID, l.expTime)

	token, err := GenerateJWT(claim, l.tokenSecret)
	if err != nil {
		return "", fmt.Errorf("error while generating token: %w", err)
	}

	err = l.LoginRepository.Insert(ctx, claim, token)
	if err != nil {
		return "", fmt.Errorf("error while inserting token: %w", err)
	}

	return token, nil
}

func GenerateJWT(claims entities.Claims, tokenSecret string) (string, error) {
	claim := jwt.RegisteredClaims{
		Subject:   claims.Sub,
		ExpiresAt: jwt.NewNumericDate(claims.ExpTime),
		IssuedAt:  jwt.NewNumericDate(claims.CreatedTime),
		ID:        claims.TokenID,
	}
	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := tokenJWT.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", fmt.Errorf("error while getting the signed token: %w", err)
	}
	return token, nil
}
