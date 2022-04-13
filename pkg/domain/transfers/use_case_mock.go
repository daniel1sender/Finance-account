package transfers

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type UseCaseMock struct {
	Transfer      entities.Transfer
	Error         error
	ListTransfers []entities.Transfer
}

func (m *UseCaseMock) Make(ctx context.Context, originID, destinationID string, amount int) (entities.Transfer, error) {
	return m.Transfer, m.Error
}

func (m *UseCaseMock) UpdateBalance(ctx context.Context, id string, balance int) error {
	panic("not implemented")
}

func (m *UseCaseMock) ListByID(ctx context.Context, accountID string) ([]entities.Transfer, error) {
	return m.ListTransfers, m.Error
}
