package accounts

import (
	"context"
	"fmt"
)

func (ar AccountRepository) DeleteAll(ctx context.Context) error {
	_, err := ar.Exec(ctx, "DELETE FROM accounts")
	if err != nil{
		return fmt.Errorf("error while deleting accounts: %w", err)
	}
	return nil
}
