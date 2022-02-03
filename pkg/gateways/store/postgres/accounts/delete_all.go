package accounts

import (
	"context"
	"fmt"
)

const deleteAllStatement = `DELETE FROM accounts`

func (ar AccountRepository) DeleteAll(ctx context.Context) error {
	_, err := ar.Exec(ctx, deleteAllStatement)
	if err != nil {
		return fmt.Errorf("error while deleting accounts: %w", err)
	}
	return nil
}
