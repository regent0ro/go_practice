package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
func CreateAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//method for Account. Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

//ChangeOwner of ther account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Balance of onwer account
func (a Account) Onwer() string {
	return a.owner
}

//override
func (a Account) String() string {
	return fmt.Sprint(a.Onwer(), "s account. Has: ", a.Balance())
}
