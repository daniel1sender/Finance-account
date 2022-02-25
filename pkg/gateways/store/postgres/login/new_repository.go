package login

import (
	pgx "github.com/jackc/pgx/v4/pgxpool"
)

type LoginRepository struct {
	*pgx.Pool
}

func NewStorage(connection *pgx.Pool) LoginRepository {
	return LoginRepository{connection}
}
