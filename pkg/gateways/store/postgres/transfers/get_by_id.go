package transfers

import (
	"context"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	"github.com/jackc/pgx/v4"
)

const getByIDStatement = "SELECT id, account_origin_id, account_destination_id, amount, created_at FROM transfers WHERE account_origin_id = $1"

func (t TransfersRepository) GetByID(ctx context.Context, accountID string) ([]entities.Transfer, error) {
	var transfer entities.Transfer
	var transfers []entities.Transfer

	rows, err := t.Query(ctx, getByIDStatement, accountID)
	if err == pgx.ErrNoRows {
		return []entities.Transfer{}, fmt.Errorf("transfers not found")
	} else if err != nil {
		return []entities.Transfer{}, err
	}

	for rows.Next() {
		err := rows.Scan(&transfer.ID, &transfer.AccountOriginID, &transfer.AccountDestinationID, &transfer.Amount, &transfer.CreatedAt)
		if err != nil {
			return []entities.Transfer{}, err
		}
		transfers = append(transfers, transfer)
	}
	if len(transfers) == 0 {
		return []entities.Transfer{}, fmt.Errorf("error while listing transfers")
	}
	return transfers, nil
}
