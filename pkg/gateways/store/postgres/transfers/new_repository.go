package transfers

import (
	pgx "github.com/jackc/pgx/v4/pgxpool"
)

type TransfersRepository struct {
	*pgx.Pool
}

func NewStorage(connection *pgx.Pool) TransfersRepository {
	return TransfersRepository{connection}
}
