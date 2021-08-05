package accounts

import (
	"fmt"

	"exemplo.com/pkg/domain/entities"
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
		if storedAccount.Cpf == cpf {
			return entities.Account{}, fmt.Errorf("account with cpf %s already exists", cpf)
		}
	}

	account, err := entities.NewAccount(name, cpf, secret, balance)

	if err != nil {
		return entities.Account{}, fmt.Errorf("err to create an new account")
	}

	au.storage[account.Id] = account

	return account, nil
}

func (au AccountUseCase) GetBalanceById(id string) (int, error) {
	for key, value := range au.storage {
		if value.Id == id {
			balance := au.storage[key].Balance
			return balance, nil
		}
	}
	return 0, fmt.Errorf("no id %s found", id)
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
			return fmt.Errorf("no account found with id %s", id)
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
		return fmt.Errorf("can't update account with value %d below zero", balance)
	}

	account.Balance = balance
	au.storage[id] = account

	return nil
}

func (au AccountUseCase) GetAccountByID(id string) (entities.Account, error) {
	account, ok := au.storage[id]
	if !ok {
		return entities.Account{}, fmt.Errorf("account with id %s not found", id)
	}

	return account, nil
}
