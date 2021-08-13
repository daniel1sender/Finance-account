package accounts

import (
	"errors"

	"exemplo.com/pkg/domain/entities"
)

var (
	ErrExistingCPF         = errors.New("cpf informed is invalid")
	ErrToCallNewAccount    = errors.New("error to call function new account")
	ErrIDNotFound          = errors.New("account id isn't found")
	ErrBalanceLessOrEqualZero = errors.New("balance account less or equal zero")
)

type AccountUseCase struct {
	storage map[string]entities.Account
}

func NewAccountUseCase(storage map[string]entities.Account) AccountUseCase {
	return AccountUseCase{
		storage: storage,
	}
}

func (au AccountUseCase) CreateAccount(name, cpf, secret string, balance int) (entities.Account, error) {

	for _, storedAccount := range au.storage {
		if storedAccount.CPF == cpf {
			return entities.Account{}, ErrExistingCPF
		}
	}

	account, err := entities.NewAccount(name, cpf, secret, balance)
	if err != nil {
		return entities.Account{}, ErrToCallNewAccount
	}

	au.storage[account.ID] = account

	return account, nil
}

func (au AccountUseCase) GetBalanceByID(id string) (int, error) {
	for key, value := range au.storage {
		if value.ID == id {
			balance := au.storage[key].Balance
			return balance, nil
		}
	}
	return 0, ErrIDNotFound
}

func (au AccountUseCase) GetAccounts() []entities.Account {
	var AccountsList []entities.Account

	for _, value := range au.storage {
		AccountsList = append(AccountsList, value)
	}

	return AccountsList
}

func (au AccountUseCase) CheckAccounts(id ...string) error {
	for _, v := range id {
		if _, ok := au.storage[v]; !ok {
			return ErrIDNotFound
		}
	}
	return nil
}

func (au AccountUseCase) UpdateAccountBalance(id string, balance int) error {
	account, err := au.GetAccountByID(id)
	if err != nil {
		return err
	}
	if balance < 0 {
		return ErrBalanceLessOrEqualZero
	}

	account.Balance = balance
	au.storage[id] = account

	return nil
}

func (au AccountUseCase) GetAccountByID(id string) (entities.Account, error) {
	account, ok := au.storage[id]
	if !ok {
		return entities.Account{}, ErrIDNotFound
	}

	return account, nil
}
