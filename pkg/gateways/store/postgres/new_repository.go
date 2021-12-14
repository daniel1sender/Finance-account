package postgres

import (
	"context"
	"log"

	pgx "github.com/jackc/pgx/v4/pgxpool"
)

type AccountRepository struct {
	*pgx.Pool
}

func NewStorage()AccountRepository{
	conn, err:= pgx.Connect(context.Background(), "postgres://postgres:1234@localhost:5432/desafio")
	if err != nil{
		log.Fatal(err)
	}
	return AccountRepository{conn}
}
