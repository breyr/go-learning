package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

// toString method
func (b Bitcoin) String() string {
	// format and return string
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// struct pointers, automatically dereferenced, don't need to do (*w).balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("oh no")
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
