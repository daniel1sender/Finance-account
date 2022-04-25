package accounts

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type Repository interface {
	GetAll(ctx context.Context) ([]entities.Account, error)
	GetBalanceByID(ctx context.Context, id string) (int, error)
	GetByID(ctx context.Context, id string) (entities.Account, error)
	CheckCPF(ctx context.Context, cpf string) error
	Upsert(ctx context.Context, account entities.Account) error
}
