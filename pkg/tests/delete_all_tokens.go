package tests

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func DeleteAllTokens(Db *pgxpool.Pool) error {
	_, err := Db.Exec(context.Background(), "DELETE FROM tokens")
	if err != nil {
		return fmt.Errorf("error while deleting tokens: %w", err)
	}
	return nil
}
