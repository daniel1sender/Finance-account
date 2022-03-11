package login

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type AccountRepository interface {
	GetByCPF(ctx context.Context, cpf string) (entities.Account, error)
}
