package loginUseCases

import (
	"context"
	"fmt"
	"time"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (l LoginUseCase) Auth(ctx context.Context, cpf, accountSecret string, duration string) (string, string, error) {
	account, err := l.AccountStorage.GetByCPF(ctx, cpf)
	if err != nil {
		return "", "", fmt.Errorf("error while getting account by cpf: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Secret), []byte(accountSecret))
	if err != nil {
		return "", "", fmt.Errorf("error while validate secret: %w", err)
	}
	expTime, err := time.ParseDuration(duration)
	if err != nil {
		return "", "", fmt.Errorf("error while parsing duration time")
	}
	tokenJWT, err := GenerateJWT(account.ID, l.tokenSecret, expTime)
	if err != nil {
		return "", "", fmt.Errorf("error while generating token: %w", err)
	}
	token, err := tokenJWT.SignedString([]byte(l.tokenSecret))
	if err != nil {
		return "", "", fmt.Errorf("error while getting the signed token: %w", err)
	}
	err = l.LoginStorage.Insert(ctx, token, l.tokenSecret)
	if err != nil {
		return "", "", fmt.Errorf("error while inserting token: %w", err)
	}

	return token, account.ID, nil
}

func GenerateJWT(accountID, tokenSecret string, expTime time.Duration) (*jwt.Token, error) {
	claim := entities.NewClaim(accountID)
	claims := jwt.RegisteredClaims{
		Subject:   accountID,
		ExpiresAt: jwt.NewNumericDate(claim.ExpTime.Add(expTime)),
		IssuedAt:  jwt.NewNumericDate(claim.CreatedTime),
		ID:        claim.TokenID,
	}
	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenJWT, nil
}

func ValidateToken(tokenString string, accountID string, tokenSecret string) error {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return fmt.Errorf("expected no error but got '%v'", err)
	}
	claims := token.Claims.(*jwt.RegisteredClaims)
	if claims.Subject != accountID {
		return fmt.Errorf("expected '%s' but got '%s'", accountID, claims.Subject)
	}
	if claims.ID == "" {
		return fmt.Errorf("expected not empty id")
	}
	if !claims.VerifyExpiresAt(time.Now(), true) {
		return fmt.Errorf("expected non-zero 'expires at' time")
	}
	if !claims.VerifyIssuedAt(time.Now(), true) {
		return fmt.Errorf("expected non-zero 'issued at' time")
	}
	return nil
}
