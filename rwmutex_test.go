package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccont struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccont) AddBalance(amount int) {
	// account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	// account.RWMutex.Unlock()
}

func (account *BankAccont) GetBalance() int {
	// account.RWMutex.RLock()
	balance := account.Balance
	// account.RWMutex.RUnlock()

	return balance
}

func TestReadWriteMutex(t *testing.T) {
	account := BankAccont{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance:", account.GetBalance())
}
