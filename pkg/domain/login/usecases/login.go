package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (l LoginUseCase) Login(ctx context.Context, cpf, accountSecret string, duration string) (string, error) {
	account, err := l.AccountStorage.GetByCPF(ctx, cpf)
	if err != nil {
		return "", fmt.Errorf("error while getting account by cpf: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Secret), []byte(accountSecret))
	if err != nil {
		return "", fmt.Errorf("error while validating secret: %w", err)
	}

	expTime, err := time.ParseDuration(duration)
	if err != nil {
		return "", fmt.Errorf("error while parsing duration time")
	}

	claim := entities.NewClaim(account.ID, expTime)
	claims := jwt.RegisteredClaims{
		Subject:   account.ID,
		ExpiresAt: jwt.NewNumericDate(claim.ExpTime),
		IssuedAt:  jwt.NewNumericDate(claim.CreatedTime),
		ID:        claim.TokenID,
	}

	token, err := GenerateJWT(claims, l.tokenSecret)
	if err != nil {
		return "", fmt.Errorf("error while generating token: %w", err)
	}

	err = l.LoginRepository.Insert(ctx, claim, token)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateJWT(claims jwt.RegisteredClaims, tokenSecret string) (string, error) {
	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenJWT.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", fmt.Errorf("error while getting the signed token: %w", err)
	}
	return token, nil
}
