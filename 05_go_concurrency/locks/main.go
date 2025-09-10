package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	balance int
	mu      sync.Mutex
}

func NewAccount(initial int) Account {
	return Account{balance: initial}
}

func (a *Account) CheckBalance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

func (a *Account) Withdraw(v int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance >= v {
		// Simulate some processing time
		time.Sleep(10 * time.Millisecond)
		a.balance -= v
		fmt.Printf("Withdrew %d, balance: %d\n", v, a.balance)
	} else {
		fmt.Printf("Insufficient funds for withdrawal of %d\n", v)
	}
}

func (a *Account) Deposit(v int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	// Simulate some processing time
	time.Sleep(10 * time.Millisecond)
	a.balance += v
	fmt.Printf("Deposited %d, balance: %d\n", v, a.balance)
}

func main() {
	account := NewAccount(100)
	var wg sync.WaitGroup // like async await in C#

	// Concurrent deposits
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			account.Deposit(amount)
		}(20)
	}

	// Concurrent withdrawals
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			account.Withdraw(amount)
		}(30)
	}

	wg.Wait()
	fmt.Printf("Final balance: %d\n", account.CheckBalance())
}
