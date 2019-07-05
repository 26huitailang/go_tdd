package pointers

import (
	"errors"
	"fmt"
)

type Bitcoint int

type Wallet struct {
	balance Bitcoint
}

func (w *Wallet) Deposit(amount Bitcoint) {
	w.balance += amount

}
func (w *Wallet) Balance() Bitcoint {

	return w.balance
}

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoint) error {
	if amount > w.balance {
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}
func (b Bitcoint) String() string {
	return fmt.Sprintf("%d BTC", b)
}
