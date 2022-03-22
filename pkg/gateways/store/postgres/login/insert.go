package login

import (
	"context"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/login"
)

const insertStatement = `INSERT INTO tokens(
	id,
	sub,
	exp_time,
	created_at,
	token
	) VALUES (
	$1, 
	$2, 
	$3,
	$4,
	$5)`

func (l LoginRepository) Insert(ctx context.Context, token, tokenSecret string) error {
	claims, err := login.ParseToken(token, tokenSecret)
	if err != nil {
		return fmt.Errorf("error while parsing token occurred: %w", err)
	}
	if _, err := l.Exec(ctx, insertStatement, claims.ID, claims.Subject, claims.ExpiresAt.Time, claims.IssuedAt.Time, token); err != nil {
		return fmt.Errorf("unable to insert the token due to: %w", err)
	}
	return nil
}
