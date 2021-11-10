package store

import "errors"

var (
	ErrExistingCPF = errors.New("cpf informed alredy exists")
)
