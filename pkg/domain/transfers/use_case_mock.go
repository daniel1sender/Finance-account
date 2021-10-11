package transfers

import "github.com/daniel1sender/Desafio-API/pkg/domain/entities"

type UseCaseMock struct {
	Transfer entities.Transfer
	Error    error
}

func (m *UseCaseMock) Make(originID, destinationID int, amount int) (entities.Transfer, error) {
	return m.Transfer, m.Error
}
