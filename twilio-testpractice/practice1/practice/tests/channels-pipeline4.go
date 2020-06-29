package main

import (
	"fmt"
	"sync"
)

func asChan(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, v := range nums {
			out <- v
		}

		close(out)
	}()

	return out
}

func Merge(channels ...<-chan int) <-chan int {

	out := make(chan int)
	wg := sync.WaitGroup{}
	output := func(ch <-chan int) {
		for v := range ch {
			out <- v
		}

		wg.Done()
	}

	wg.Add(len(channels))

	for _, ch := range channels {
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	//closing output channel

	return out

}

func Merge2(a, b <-chan int) <-chan int {

	out := make(chan int)

	go func() {

		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					fmt.Println("a is done")
					a = nil
					continue
				}

				out <- v

			case v, ok := <-b:
				if !ok {
					fmt.Println("b is done")
					b = nil
					continue
				}

				out <- v
			}
		}

		close(out)
	}()

	return out

}

func readChan(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	a := asChan(1, 2, 3, 4, 5)
	b := asChan(6, 7, 8, 9, 10)

	c := Merge(a, b)
	readChan(c)

}
