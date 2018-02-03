package account

import (
	"sync"
)

var mutex = &sync.Mutex{}

// Account struct
type Account struct {
	balance  int64
	isClosed bool
}

// Open a new account with initial deposit
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit, isClosed: false}
}

// Balance returns current balance of the account
func (a *Account) Balance() (int64, bool) {
	if a.isClosed {
		return 0, false
	}
	return a.balance, true
}

// Close balance and no money can deposit to the account
func (a *Account) Close() (int64, bool) {
	mutex.Lock()
	balance, flag := a.balance, true
	if a.isClosed == false {
		a.isClosed = true
	} else {
		balance, flag = 0, false
	}
	mutex.Unlock()
	return balance, flag
}

// Deposit deposits amount to account to update to new balance
func (a *Account) Deposit(amt int64) (int64, bool) {
	mutex.Lock()
	newBal, ok := a.balance+amt, false
	if newBal >= 0 && a.isClosed == false {
		a.balance, ok = newBal, true
	}
	mutex.Unlock()
	return a.balance, ok
}
