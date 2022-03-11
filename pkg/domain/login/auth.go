package login

import (
	"context"
	"fmt"
	"time"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (l LoginUseCase) Auth(ctx context.Context, cpf, accountSecret string, duration string) (string, error) {
	account, err := l.AccountStorage.GetByCPF(ctx, cpf)
	if err != nil {
		return "", fmt.Errorf("error while getting account by cpf: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Secret), []byte(accountSecret))
	if err != nil {
		return "", fmt.Errorf("error while validate secret: %w", err)
	}

	expTime, err := time.ParseDuration(duration)
	if err != nil {
		return "", fmt.Errorf("error while parsing duration time")
	}

	token, err := GenerateJWT(account.ID, l.tokenSecret, expTime)
	if err != nil {
		return "", fmt.Errorf("error while generating token: %w", err)
	}
	return token, nil
}

func GenerateJWT(accountID, tokenSecret string, expTime time.Duration) (string, error) {
	claim := entities.NewClaim(accountID)
	claims := jwt.RegisteredClaims{
		Subject:   accountID,
		ExpiresAt: jwt.NewNumericDate(claim.ExpTime.Add(expTime)),
		IssuedAt:  jwt.NewNumericDate(claim.CreatedTime),
		ID:        claim.TokenID,
	}
	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenJWT.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", fmt.Errorf("error while getting the signed token: %w", err)
	}
	return token, nil
}