package login

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/login"
	"github.com/jackc/pgx/v4"
)

func (l LoginRepository) CheckToken(ctx context.Context, token string) error {
	var tokenString string
	err := l.QueryRow(ctx, "SELECT token FROM tokens WHERE token = $1", token).Scan(&tokenString)
	if err == pgx.ErrNoRows {
		return login.ErrTokenNotFound
	} else if err != nil {
		return err
	}
	return nil
}
