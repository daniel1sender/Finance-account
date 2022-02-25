package login

import (
	"context"
	"fmt"

	jwt "github.com/golang-jwt/jwt/v4"
)

const insertStatement = `INSERT INTO tokens(
	id,
	token,
	created_at
	) VALUES (
	$1, 
	$2, 
	$3)`

func (l LoginRepository) Insert(ctx context.Context, token, tokenSecret string) error {
	tokenParsed, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return fmt.Errorf("error while parsing token occurred %w", err)
	}
	claims := tokenParsed.Claims.(*jwt.RegisteredClaims)
	if _, err := l.Exec(ctx, insertStatement, claims.ID, token, claims.IssuedAt.Time); err != nil {
		return fmt.Errorf("unable to insert the token due to: %v", err)
	}
	return nil
}
