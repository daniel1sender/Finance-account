package domain

import "fmt"

type Transfer struct {
	Id                      string
	Account_origin_id       int
	Account_destinantion_id int
	Amount                  float64
	//Created_at time.Time //?
}

var transfersMap = make(map[int]Transfer)
var transferNumber = 0
//Essa função é similar a função createAccount

func MakeTransfer(t Transfer)(map[int]Transfer, error){
	count := 0


	if t.Amount == 0 {
		return nil, fmt.Errorf("amount equal zero")
	}

	if t.Account_origin_id == t.Account_destinantion_id{
		return nil, fmt.Errorf("transfer is to the same id")
	}

	for _, v := range AccountsMap{
		if t.Account_origin_id == v.Id{
			count++
		}
		if t.Account_destinantion_id == v.Id{
			count++
		}
	}

	if count == 2{
		transfersMap[transferNumber] = t
		transferNumber++
		return transfersMap, nil 
	}

	return nil, fmt.Errorf("id not found")
}