package transfers

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func DeleteAll(Db *pgxpool.Pool) error {
	_, err := Db.Exec(context.Background(), "DELETE FROM transfers")
	if err != nil {
		return fmt.Errorf("error while deleting table transfers: %w", err)
	}
	return nil
}
