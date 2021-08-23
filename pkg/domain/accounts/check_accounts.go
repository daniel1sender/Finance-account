package accounts

func (au AccountUseCase) CheckAccounts(id ...string) error {
	return au.storage.CheckAccounts(id)
}
