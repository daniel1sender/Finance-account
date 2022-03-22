package login

import (
	"context"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
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

func (l LoginRepository) Insert(ctx context.Context, claims entities.Claims, token string) error {
	if _, err := l.Exec(ctx, insertStatement, claims.TokenID, claims.Sub, claims.ExpTime, claims.CreatedTime, token); err != nil {
		return fmt.Errorf("unable to insert the token due to: %w", err)
	}
	return nil
}
