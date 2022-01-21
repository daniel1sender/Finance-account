package usecases

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func DeleteAll(Db *pgxpool.Pool) error {
	_, err := Db.Exec(context.Background(), "DELETE FROM accounts")
	if err != nil {
		return fmt.Errorf("error while deleting accounts: %w", err)
	}
	return nil
}
