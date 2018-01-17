package account

import (
	"fmt"
	"sync"
)

// Personal ...
type Personal struct {
	amount float32
	// lock is needed to ensure only a single routine can access the values within
	lock sync.Mutex
}

func CreateAccount() *Personal {
	return &Personal{
		amount: 0.0,
		lock:   sync.Mutex{},
	}
}

func (a *Personal) Withdraw(amount float32) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	switch {
	case a.amount-amount >= 0.0:
		// All good
		a.amount -= amount
	default:
		// Account is overdrawn
		return fmt.Errorf("The account is overdrawn")
	}
	return nil
}

func (a *Personal) Deposit(amount float32) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.amount += amount
}

func (a *Personal) Amount() float32 {
	return a.amount
}
