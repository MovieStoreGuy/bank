package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/MovieStoreGuy/bank/stateful/account"
)

// Statefull implementation of a bank
func main() {
	wg := &sync.WaitGroup{}
	ac := account.CreateAccount()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(account *account.Personal, id int, w *sync.WaitGroup) {
			// Sends channel updates based on tick intervals
			inc := time.Tick(10 * time.Millisecond)
			dec := time.Tick(15 * time.Millisecond)
			// Send channel updates once time period has finished
			done := time.After(10 * time.Second)
			defer w.Done()
			for {
				select {
				case <-inc:
					ac.Deposit(10.37)
				case <-dec:
					if err := ac.Withdraw(15.82); err != nil {
					}
				case <-done:
					return
				}
				fmt.Printf("Thread %2d has updated the value to %.2f\r", id, account.Amount())
			}
		}(ac, i, wg)
	}
	wg.Wait()
	fmt.Printf("\nThe final amount is: %.2f\n", ac.Amount())
}
