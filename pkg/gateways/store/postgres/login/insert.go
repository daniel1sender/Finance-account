package login

import (
	"context"
	"fmt"
	"time"
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

func (l LoginRepository) Insert(ctx context.Context, tokenID, accountID, token string, expiresAt, createdAt time.Time) error {
	if _, err := l.Exec(ctx, insertStatement, tokenID, accountID, expiresAt, createdAt, token); err != nil {
		return fmt.Errorf("unable to insert the token due to: %w", err)
	}
	return nil
}
