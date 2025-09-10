package main

import (
	"fmt"
	"sync"
	"time"
)

type AccountWithRW struct {
	balance int
	rw      sync.RWMutex
}

func NewAccountWithRW(init int) *AccountWithRW {
	return &AccountWithRW{balance: init}
}

func (a *AccountWithRW) CheckBalance() int {
	a.rw.RLock()
	defer a.rw.RUnlock()

	time.Sleep(20 * time.Millisecond) // Simulate some processing time
	return a.balance
}

func (a *AccountWithRW) IsPositive() bool {
	a.rw.RLock()
	defer a.rw.RUnlock()

	time.Sleep(10 * time.Millisecond) // Simulate some processing time
	return a.balance >= 0
}

func (a *AccountWithRW) Deposit(v int) {
	a.rw.Lock()
	defer a.rw.Unlock()

	time.Sleep(10 * time.Millisecond) // Simulate some processing time
	a.balance += v
	fmt.Printf("Deposited %d, balance: %d\n", v, a.balance)
}

func (a *AccountWithRW) Withdraw(v int) {
	a.rw.Lock()         // Writer lock (exclusive)
	defer a.rw.Unlock() // Writer unlock
	if a.balance >= v {
		time.Sleep(10 * time.Millisecond)
		a.balance -= v
		fmt.Printf("Withdrew %d, balance: %d\n", v, a.balance)
	}
}

func main() {
	account := NewAccountWithRW(100)

	wg := sync.WaitGroup{}

	// many readers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			balance := account.CheckBalance() // RLock - không block nhau
			fmt.Printf("Reader %d: balance = %d\n", id, balance)
		}(i)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		account.Deposit(50) // Lock - block tất cả readers và writers khác
	}()

	wg.Wait()
	fmt.Printf("Final balance: %d\n", account.CheckBalance())
}
