package usecases

import (
	"context"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/daniel1sender/Desafio-API/pkg/domain/login"
)

func (l LoginUseCase) ValidateToken(ctx context.Context, tokenString string) (entities.Claims, error) {
	claim, err := l.ParseToken(ctx, tokenString)
	if err != nil {
		return entities.Claims{}, fmt.Errorf("%w: %v", login.ErrInvalidToken, err)
	}
	token, err := l.LoginRepository.GetTokenByID(ctx, claim.TokenID)
	if err != nil {
		return entities.Claims{}, fmt.Errorf("error while getting token: %w", err)
	}
	if tokenString != token {
		return entities.Claims{}, fmt.Errorf("error while comparing token: %w", login.ErrInvalidToken)
	}
	return claim, nil
}
