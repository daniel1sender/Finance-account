package login

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

var (
	ErrTokenNotFound = errors.New("token not found")
)

func (l LoginRepository) CheckToken(ctx context.Context, token string) error {
	var tokenString string
	err := l.QueryRow(ctx, "SELECT token FROM tokens WHERE token = $1", token).Scan(&tokenString)
	if err == pgx.ErrNoRows {
		return ErrTokenNotFound
	} else if err != nil {
		return err
	}
	return nil
}
