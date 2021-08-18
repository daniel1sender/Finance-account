package accounts

func (au AccountUseCase) GetBalanceByID(id string) (int, error) {
	return au.storage.FindBalanceByID(id)

}
