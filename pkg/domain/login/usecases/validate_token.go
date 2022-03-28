package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

var (
	ErrInvalidToken = errors.New("invalid token found")
)

func (l LoginUseCase) ValidateToken(ctx context.Context, tokenString string) (entities.Claims, error) {
	claim, err := l.ParseToken(ctx, tokenString)
	if err != nil {
		return entities.Claims{}, fmt.Errorf("error while validating token: %w", err)
	}
	token, err := l.LoginRepository.GetTokenByID(ctx, claim.TokenID)
	if err != nil {
		return entities.Claims{}, fmt.Errorf("error while getting token: %w", err)
	}
	if tokenString != token {
		return entities.Claims{}, fmt.Errorf("error while comparing token :%w", ErrInvalidToken)
	}
	return claim, nil
}
