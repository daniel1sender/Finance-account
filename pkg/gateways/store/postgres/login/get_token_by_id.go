package login

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func (l LoginRepository) GetTokenByID(ctx context.Context, tokenID string) (string, error) {
	var token string
	err := l.QueryRow(ctx, "SELECT token FROM tokens WHERE id = $1", tokenID).Scan(&token)
	if err == pgx.ErrNoRows {
		return "", ErrTokenNotFound
	} else if err != nil {
		return "", err
	}
	return token, nil
}
