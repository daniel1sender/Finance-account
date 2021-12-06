package usecases

func (au AccountUseCase) GetBalanceByID(id string) (int, error) {
	balance, err := au.storage.GetBalanceByID(id)
	return balance, err
}
