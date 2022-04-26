package accounts

import (
	pgx "github.com/jackc/pgx/v4/pgxpool"
)

type AccountRepository struct {
	*pgx.Pool
}

func NewRepository(connection *pgx.Pool) AccountRepository {
	return AccountRepository{connection}
}
