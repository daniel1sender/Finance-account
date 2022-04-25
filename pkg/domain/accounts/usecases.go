package accounts

import (
	"context"
	"errors"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

var (
	ErrExistingCPF     = errors.New("cpf informed alredy exists")
	ErrAccountNotFound = errors.New("account not found")
	ErrEmptyList       = errors.New("empty list of accounts")
)

type UseCase interface {
	GetBalanceByID(ctx context.Context, id string) (int, error)
	Create(ctx context.Context, name, cpf, secret string, balance int) (entities.Account, error)
	GetByID(ctx context.Context, id string) (entities.Account, error)
	GetAll(ctx context.Context) ([]entities.Account, error)
}
