package domain

import (
	"errors"
)

var (
	ErrEmptySecret = errors.New("empty secret was informed")
	ErrInvalidCPF  = errors.New("cpf informed is invalid")
)

func ValidateCPF(cpf string) error {
	if len(cpf) != 11 {
		return ErrInvalidCPF
	}
	return nil
}

func ValidateSecret(secret string) error {
	if secret == "" {
		return ErrEmptySecret
	}
	return nil
}
