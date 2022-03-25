package usecases

import (
	"context"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	jwt "github.com/golang-jwt/jwt/v4"
)

func (l LoginUseCase) ValidateToken(ctx context.Context, tokenString string) (entities.Claims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("invalid signature method")
		}
		return []byte(l.tokenSecret), nil
	}
	tokenParsed, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, keyFunc)
	if err != nil {
		return entities.Claims{}, fmt.Errorf("error while validating token: %w", err)
	}
	claims := tokenParsed.Claims.(*jwt.RegisteredClaims)
	token := entities.Claims{
		TokenID:     claims.ID,
		Sub:         claims.Subject,
		ExpTime:     claims.ExpiresAt.Time,
		CreatedTime: claims.IssuedAt.Time,
	}
	_, err = l.LoginRepository.GetTokenByID(ctx, token.TokenID)
	if err != nil {
		return entities.Claims{}, fmt.Errorf("error while validating token: %w", err)
	}
	return token, nil
}
