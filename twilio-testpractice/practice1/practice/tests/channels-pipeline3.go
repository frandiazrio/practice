package main

import (
	"fmt"
	"sync"
)

func asChan(vs ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, v := range vs {
			out <- v
		}

		close(out)
	}()

	return out
}

func merge2(a, b <-chan int) <-chan int {
	wg := sync.WaitGroup{}
	out := make(chan int)

	output := func(ch <-chan int) {

		for v := range ch {
			out <- v
		}

		wg.Done()

	}

	wg.Add(2)

	go output(a)
	go output(b)

	go func() {
		wg.Wait()
		close(out)
	}()
	wg.Wait()
	return out

}

/*
func merge(a, b <-chan int) <-chan int {

	out := make(chan int)

	go func() {
		defer close(out)

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

	}()

	return out
}*/

func read(in <-chan int) {

	for num := range in {
		fmt.Println(num)
	}

}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)

	m := merge2(a, b)

	read(m)

}
