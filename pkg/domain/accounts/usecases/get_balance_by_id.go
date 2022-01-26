package usecases

func (au AccountUseCase) GetBalanceByID(id string) (int, error) {
	balance, err := au.storage.GetBalanceByID(id)
	if err != nil {
		return balance, err
	}
	return balance, nil
}
