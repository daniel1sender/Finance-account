package transfers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (tr TransferRepository) UpdateByID(transfer entities.Transfer) error {
	tr.transfers[transfer.ID] = entities.Transfer{ID: transfer.ID, AccountOriginID: transfer.AccountOriginID, AccountDestinationID: transfer.AccountDestinationID}
	keepTransfer, err := json.MarshalIndent(tr.transfers, "", " ")
	if err != nil {
		return fmt.Errorf("error decoding account '%v'", err)
	}
	_ = os.WriteFile(tr.storage.Name(), keepTransfer, 0644)
	return nil
}
