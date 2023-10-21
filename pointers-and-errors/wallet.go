package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(b Bitcoin) {
	w.balance += b
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrNotEnoughFunds = errors.New("Not enough minerals.")

func (w *Wallet) Withdraw(b Bitcoin) error {
	if b > w.balance {
		return ErrNotEnoughFunds
	}
	w.balance -= b
	return nil

}
