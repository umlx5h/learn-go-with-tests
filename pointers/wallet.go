package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// var ErrInsufficientFunds = "cannot withdraw, insufficient funds: "
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		// return fmt.Errorf(ErrInsufficientFunds+"%v", amount)
		return ErrInsufficientFunds
		// return fmt.Errorf("unko")
	}

	w.Deposit(-amount)

	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
