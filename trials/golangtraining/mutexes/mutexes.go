package main

import (
	"fmt"
	"sync"
)

var (
	mu      sync.Mutex // guards balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	// If you uncomment the below statement, you will see the
	// 	update happening as expected and in certain order.
	// fmt.Println("Current balance inside: ", balance)
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

func main() {
	// Deposit [1..1000] concurrently.

	max := 100
	var n sync.WaitGroup
	for i := 1; i <= max; i++ {
		n.Add(1)
		go func(amount int) {
			Deposit(amount)
			// If you uncomment the below statement, you will see the
			//	that the actual reading of the data is uncertain.
			// fmt.Println("Current balance outside: ", Balance())
			n.Done()
		}(i)
	}
	n.Wait()

	got, want := Balance(), (max+1)*max/2
	if got != want {
		fmt.Errorf("Balance = %d, want %d", got, want)
	} else {
		fmt.Printf("All Ok! Balance = %d, want %d\n", got, want)
	}
}
