package main

import (
	"fmt"
	"sync"
	"time"
)

type DepositRequest struct {
	Amount int
}

type WithdrawRequest struct {
	Amount int
	Result chan bool // Channel để trả kết quả
}

type BalanceRequest struct {
	Result chan int // Channel để trả balance
}

type AccountChannel struct {
	depositCh  chan DepositRequest
	withdrawCh chan WithdrawRequest
	balanceCh  chan BalanceRequest
	doneCh     chan bool
}

func NewAccountChannel(initial int) *AccountChannel {
	acc := &AccountChannel{
		depositCh:  make(chan DepositRequest),
		withdrawCh: make(chan WithdrawRequest),
		balanceCh:  make(chan BalanceRequest),
		doneCh:     make(chan bool),
	}

	go func() {
		balance := initial // ← Chỉ có goroutine này được touch balance

		for {
			select {
			case req := <-acc.depositCh:
				time.Sleep(10 * time.Millisecond) // Simulate some processing time
				balance += req.Amount             // Safe
				fmt.Printf("Deposited %d, balance: %d\n", req.Amount, balance)
			case req := <-acc.balanceCh:
				time.Sleep(20 * time.Millisecond) // Simulate some processing time
				req.Result <- balance             // Safe
			case req := <-acc.withdrawCh:
				time.Sleep(10 * time.Millisecond) // Simulate some processing time

				if balance >= req.Amount {
					balance -= req.Amount
					fmt.Printf("Withdrew %d, balance: %d\n", req.Amount, balance)
					req.Result <- true // ← Gửi kết quả về client
				} else {
					fmt.Printf("Insufficient funds for withdrawal of %d\n", req.Amount)
					req.Result <- false // ← Gửi kết quả về client
				}
			case <-acc.doneCh:
				fmt.Println("Account goroutine exiting...")
				return
			}
		}
	}()

	return acc
}

func (a *AccountChannel) Deposit(v int) {
	a.depositCh <- DepositRequest{Amount: v}
}

func (a *AccountChannel) Withdraw(v int) bool {
	result := make(chan bool)
	a.withdrawCh <- WithdrawRequest{Amount: v, Result: result}
	return <-result // Wait for result
}

func (a *AccountChannel) CheckBalance() int {
	result := make(chan int)
	a.balanceCh <- BalanceRequest{Result: result}
	return <-result
}

func (a *AccountChannel) Close() {
	a.doneCh <- true
}

func main() {
	fmt.Println("=== Bank Account with Channels ===")
	account := NewAccountChannel(100)
	defer account.Close()

	var wg sync.WaitGroup

	// Multiple readers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			balance := account.CheckBalance() // ← Các calls này VẪN tuần tự!
			fmt.Printf("Reader %d: balance = %d\n", id, balance)
		}(i)
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			account.Deposit(amount)
		}(20)
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			success := account.Withdraw(amount)
			if !success {
				fmt.Printf("Failed to withdraw %d due to insufficient funds\n", amount)
			} else {
				fmt.Printf("Successfully withdrew %d\n", amount)
			}
		}(30)
	}

	wg.Wait()
	fmt.Printf("Final balance: %d\n", account.CheckBalance())
}
