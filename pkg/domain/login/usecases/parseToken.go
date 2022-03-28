package usecases

import (
	"context"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	jwt "github.com/golang-jwt/jwt/v4"
)

func (l LoginUseCase) ParseToken(ctx context.Context, tokenString string) (entities.Claims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("invalid signature method")
		}
		return []byte(l.tokenSecret), nil
	}
	tokenParsed, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, keyFunc)
	if err != nil {
		return entities.Claims{}, fmt.Errorf("error parsing token: %w", err)
	}
	claim := tokenParsed.Claims.(*jwt.RegisteredClaims)
	claims := entities.Claims{
		TokenID:     claim.ID,
		Sub:         claim.Subject,
		ExpTime:     claim.ExpiresAt.Time,
		CreatedTime: claim.IssuedAt.Time,
	}
	return claims, nil
}
