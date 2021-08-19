package accounts

func (s AccountStorage) CheckAccountsByID(id []string) error {
	for _, v := range id {
		if _, ok := s.storage[v]; !ok {
			return ErrIDNotFound
		}
	}
	return nil
}
