package transfers

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
)

func (tr TransferRepository) UpdateByID(transfer entities.Transfer) {
	tr.users[transfer.ID] = entities.Transfer{ID: transfer.ID, AccountOriginID: transfer.AccountOriginID, AccountDestinationID: transfer.AccountDestinationID}
	keepTransfer, err := json.MarshalIndent(tr.users, "", " ")
	if err != nil {
		log.Fatal("error decoding account")
	}
	_ = ioutil.WriteFile("Transfer_Respository.json", keepTransfer, 0644)
}
