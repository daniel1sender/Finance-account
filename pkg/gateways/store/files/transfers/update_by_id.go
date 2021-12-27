package transfers

import (
	"encoding/json"
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (tr TransferRepository) UpdateByID(transfer entities.Transfer) error {
	tr.transfers[transfer.ID] = entities.Transfer{ID: transfer.ID, AccountOriginID: transfer.AccountOriginID, AccountDestinationID: transfer.AccountDestinationID}
	keepTransfer, err := json.MarshalIndent(tr.transfers, "", " ")
	if err != nil {
		return fmt.Errorf("error decoding account: %v", err)		
	}
	_, err = tr.storage.Write(keepTransfer)
	if err != nil {
		return fmt.Errorf("error while writing in file: %v", err)
	}
	return nil
}
