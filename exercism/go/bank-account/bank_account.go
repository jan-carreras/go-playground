package account

import (
	"sync"
)

type Account struct {
	amount int
	active bool
	mux    sync.RWMutex
}

func Open(amount int) *Account{
	if amount < 0 {
		return nil
	}

	return &Account{
		amount: amount,
		active: true,
		mux: sync.RWMutex{},
	}
}

func (a *Account) Balance() (balance int, ok bool) {
	a.mux.RLock()
	defer a.mux.RUnlock()

	return a.amount, a.active
}

func (a *Account) Close() (payout int, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if !a.active {
		return a.amount, false
	}

	payout = a.amount
	a.active = false
	a.amount = 0

	return payout, true
}

func (a *Account) Deposit(amount int) (newBalance int, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if !a.active {
		return a.amount, false
	}

	newBalance = a.amount + amount
	if amount < 0 && newBalance<0 {
		// Not enough funds
		return a.amount, false
	}
	a.amount = newBalance

	return newBalance, true
}
