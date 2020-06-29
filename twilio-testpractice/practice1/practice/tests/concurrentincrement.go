package main

import (
	"fmt"
	"sync"
)

var accountBalance = 0
var mux = &sync.RWMutex{}
var wg = &sync.WaitGroup{}

func updateBalance(amt int, wg *sync.WaitGroup) {
	defer wg.Done()
	mux.Lock()
	fmt.Println(accountBalance)
	accountBalance += amt

	mux.Unlock()
}

func main() {

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			updateBalance(1, wg)
		}()

	}

	wg.Wait()
}
