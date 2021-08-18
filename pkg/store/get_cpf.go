package store

func (s AccountStorage) GetCPF(cpf string) error {
	for _, storedAccount := range s.storage {
		if storedAccount.CPF == cpf {
			return ErrExistingCPF
		}
	}
	return nil
}
