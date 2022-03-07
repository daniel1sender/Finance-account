package login

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/accounts"
	"github.com/jackc/pgx/v4"
)

func (l LoginRepository) GetTokenByID(ctx context.Context, id string) (string, error) {
	var token string
	err := l.QueryRow(ctx, "SELECT token FROM tokens WHERE sub = $1", id).Scan(&token)
	if err == pgx.ErrNoRows {
		return "", accounts.ErrAccountNotFound
	} else if err != nil {
		return "", err
	}
	return token, nil
}
