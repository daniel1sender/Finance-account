package transfers

import (
	"context"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

const updateByIDStatement = `INSERT INTO transfers(
	id,
	account_origin_id,
	account_destination_id,
	amount, created_at
	) VALUES (
	$1, 
	$2, 
	$3, 
	$4, 
	$5)`

func (tr TransfersRepository) Insert(ctx context.Context, transfer entities.Transfer) error {
	if _, err := tr.Exec(ctx, updateByIDStatement, transfer.ID, transfer.AccountOriginID, transfer.AccountDestinationID, transfer.Amount, transfer.CreatedAt); err != nil {
		return fmt.Errorf("unable to insert the transfer due to: %v", err)
	}
	return nil
}
