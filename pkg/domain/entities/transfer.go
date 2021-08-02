package entities

type Transfer struct {
	Id                      int
	Account_origin_id       int
	Account_destinantion_id int
	Amount                  float64
	//Created_at time.Time //?
}
