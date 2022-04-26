package login

import (
	pgx "github.com/jackc/pgx/v4/pgxpool"
)

type LoginRepository struct {
	*pgx.Pool
}

func NewRepository(connection *pgx.Pool) LoginRepository {
	return LoginRepository{connection}
}
