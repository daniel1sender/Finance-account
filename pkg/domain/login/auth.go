package login

import (
	"context"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (l LoginUseCase) Auth(ctx context.Context, cpf, secret string) (string, error) {
	account, err := l.AccountStorage.GetByCPF(ctx, cpf)
	if err != nil {
		return "", fmt.Errorf("error while getting account by cpf: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Secret), []byte(secret))
	if err != nil {
		return "", fmt.Errorf("error while validate secret: %w", err)
	}

	tokenJWT, err := GenerateJWT(account.ID, l.tokenSecret)
	if err != nil {
		return "", fmt.Errorf("error while generating token: %w", err)
	}

	token, err := tokenJWT.SignedString([]byte(l.tokenSecret))
	if err != nil {
		return "", fmt.Errorf("error while getting the signed token: %w", err)
	}

	err = l.LoginStorage.Insert(ctx, token, l.tokenSecret)
	if err != nil {
		return "", fmt.Errorf("error while inserting token: %w", err)
	}
	return token, nil
}

func GenerateJWT(accountID string, tokenSecret string) (*jwt.Token, error) {
	claim := entities.NewClaim(accountID)
	claims := jwt.RegisteredClaims{
		Subject:   accountID,
		ExpiresAt: jwt.NewNumericDate(claim.ExpTime),
		IssuedAt:  jwt.NewNumericDate(claim.CreatedTime),
		ID:        claim.TokenID,
	}
	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tokenJWT, nil
}
