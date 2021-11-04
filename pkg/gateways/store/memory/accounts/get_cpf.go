package accounts

func (s AccountStorage) CheckCPF(cpf string) error {
	for _, storedAccount := range s.storage {
		if storedAccount.CPF == cpf {
			return ErrExistingCPF
		}
	}
	return nil
}
