package usecases

import (
	"context"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/domain/login"
	"github.com/daniel1sender/Desafio-API/pkg/domain/verify"
	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (l LoginUseCase) Login(ctx context.Context, cpf, accountSecret string) (string, error) {
	err := verify.IsValidSecret(accountSecret)
	if err != nil {
		return "", fmt.Errorf("error while validating the secret informed: %w", verify.ErrEmptySecret)
	}
	err = verify.IsValidCPF(cpf)
	if err != nil {
		return "", fmt.Errorf("error while validating the cpf informed: %w", verify.ErrInvalidCPF)
	}
	account, err := l.AccountStorage.GetByCPF(ctx, cpf)
	if err != nil {
		return "", fmt.Errorf("error while getting the account by cpf: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Secret), []byte(accountSecret))
	if err != nil {
		return "", fmt.Errorf("error while comparing the secret informed with stored: %w", login.ErrInvalidSecret)
	}

	claim := entities.NewClaim(account.ID, l.expTime)

	token, err := GenerateJWT(claim, l.tokenSecret)
	if err != nil {
		return "", fmt.Errorf("error while generating the token: %w", err)
	}

	err = l.LoginRepository.Insert(ctx, claim, token)
	if err != nil {
		return "", fmt.Errorf("error while inserting the token: %w", err)
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
