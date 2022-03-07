package login

import (
	"context"
	"fmt"

	jwt "github.com/golang-jwt/jwt/v4"
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
	tokenParsed, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return fmt.Errorf("error while parsing token occurred: %w", err)
	}
	claims := tokenParsed.Claims.(*jwt.RegisteredClaims)
	if _, err := l.Exec(ctx, insertStatement, claims.ID, claims.Subject, claims.ExpiresAt.Time, claims.IssuedAt.Time, token); err != nil {
		return fmt.Errorf("unable to insert the token due to: %v", err)
	}
	return nil
}