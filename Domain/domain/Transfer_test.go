package domain_test

import (
	"testing"

	"exemplo.com/domain"
)
func TestMakeTransfer(t *testing.T) {

	t.Run("Should return transfersMap when amount is diferrent from zero", func(t *testing.T){
		transfer := domain.Transfer{
			Id:"1",
			Account_origin_id:2 ,
			Account_destinantion_id:3,
			Amount:10,
		}

 		account1 := domain.Account{
			Id:      2,
			Name:    "joão",
			Cpf:     "12345678910",
			Balance: 10,
		}

		account2 := domain.Account{
			Id:      3,
			Name:    "joão",
			Cpf:     "12345678910",
			Balance: 10,
		}
		 
		domain.AccountsMap[1] = account1
		domain.AccountsMap[2] = account2 

		transferMap, err := domain.MakeTransfer(transfer)

		if len(transferMap) == 0 {
			t.Errorf("transferMap should not be empty %v", transferMap)
		}

		if err != nil{
			t.Errorf("err should be nil but it is %s", err)
		}

	})

	t.Run("If amount equal zero the trasfer is not created", func(t *testing.T){
		transfer := domain.Transfer{
			Id:"1",
			Account_origin_id:2 ,
			Account_destinantion_id:3,
			Amount:0,
		}

		expectedErrMessage := "amount equal zero"

		transfersMap, err := domain.MakeTransfer(transfer)

		if err.Error() != expectedErrMessage{
			t.Errorf("err must be %s but it is %s", expectedErrMessage, err)
		}

		if transfersMap != nil {
			t.Errorf("transferMap must be nil but it is %v", transfersMap)
		}

	})

	t.Run("should return a map and a message error when the transfer is to the different accounts\n", func(t *testing.T){
		transfer := domain.Transfer{
			Id:"1",
			Account_origin_id:3 ,
			Account_destinantion_id:2,
			Amount:10,
		}

		transferMap, err:= domain.MakeTransfer(transfer)

		if len(transferMap) == 0 {
			t.Errorf("transferMap must be %v but it is nil", transferMap)
		}

		if err != nil {
			t.Errorf("err must be nil but it is %s", err)
		}

	})

	t.Run("should return a empty map and a message error when the transfer is to the same account\n", func(t *testing.T) {
		transfer := domain.Transfer{
			Id:"1",
			Account_origin_id:2 ,
			Account_destinantion_id:2,
			Amount:10,
		}

		expectedErrMessage := "transfer is to the same id"

		transferMap, err :=  domain.MakeTransfer(transfer)

		if transferMap != nil {
			t.Errorf("transferMap must be nil but it is %v", transferMap)
		}

		if err.Error() != expectedErrMessage{
			t.Errorf("Err must be %s but it is %s", expectedErrMessage, err)
		}
	}) 

	t.Run("should return a empty map and a error message when any of ids are not found", func(t *testing.T){

		transfer := domain.Transfer{
			Id:"1",
			Account_origin_id:6 ,
			Account_destinantion_id:4,
			Amount:10,
		}
		
		expectedErrMessage := "id not found"
		
		transferMap, err := domain.MakeTransfer(transfer)

		if transferMap != nil{
			t.Errorf("transferMap should be nil but it is %v", transferMap)
		}

		if err!=nil && err.Error() != expectedErrMessage{
			t.Errorf("err menssage should be %s but it is %s", expectedErrMessage, err)
		}

	})

	t.Run("should return a map when the ids are from different accounts", func(t *testing.T){
		transfer := domain.Transfer{
			Id:"1",
			Account_origin_id:3 ,
			Account_destinantion_id:2,
			Amount:10,
		}

		transferMap, err := domain.MakeTransfer(transfer)

		if len(transferMap) == 0 {
			t.Errorf("transferMap should not be empty %v", transferMap)
		}

		if err != nil{
			t.Errorf("err should be nil but it is %s", err)
		}
	})


}
