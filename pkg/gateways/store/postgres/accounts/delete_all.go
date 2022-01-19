package accounts

import (
	"context"
	"fmt"
)

func (ar AccountRepository) DeleteAll() error {
	_, err := ar.Exec(context.Background(), "DELETE FROM accounts")
	if err != nil{
		return fmt.Errorf("error while deleting accounts: %w", err)
	}
	return nil
}
