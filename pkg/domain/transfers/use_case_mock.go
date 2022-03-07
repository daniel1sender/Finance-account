package transfers

import (
	"context"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

type UseCaseMock struct {
	Transfer entities.Transfer
	Error    error
}

func (m *UseCaseMock) Make(ctx context.Context, originID, destinationID string, amount int) (entities.Transfer, error) {
	return m.Transfer, m.Error
}

func (m *UseCaseMock) UpdateBalance(ctx context.Context, id string, balance int) error {
	panic("not implemented")
}

func (m *UseCaseMock) GetByID(ctx context.Context, accountID, token, tokenSecret string) ([]entities.Transfer, error) {
	panic("not implemented")
}
