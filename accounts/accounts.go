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

var errNoMoney = errors.New(
	"Can't withdraw you are poor",
)

func NewAccout(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x ammount on your account
func (a *Account) Deposit(ammount int) { //receiver 자바스크립트에서 객체 안에 함수를 넣는 대신에 사용하는 거라고 할 수 있음
	a.balance += ammount
}

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

// Change Owner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner
func (a *Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas : ", a.Balance())
}
