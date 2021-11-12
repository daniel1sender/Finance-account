package store

import "errors"

var (
	ErrExistingCPF = errors.New("cpf informed alredy exists")
	ErrIDNotFound  = errors.New("account id isn't found")
)
