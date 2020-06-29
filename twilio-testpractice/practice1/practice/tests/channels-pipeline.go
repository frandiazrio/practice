package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
	ch := make(chan int)

	go func() {
		for _, i := range nums {
			ch <- i
		}

		close(ch)
	}()

	return ch
}

func square(input <-chan int) <-chan int {
	output := make(chan int)

	go func() {
		for v := range input {
			output <- v * v
		}

		close(output)
	}()

	return output

}

func merge(chs ...<-chan int) <-chan int {
	wg := sync.WaitGroup{}

	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(chs))

	for _, c := range chs {

		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out

}

func main() {

	numbers := make([]int, 10000000)

	for i := 0; i < len(numbers); i++ {
		numbers[i] = i

	}

	naturals := gen(numbers...)

	res1 := square(naturals)
	res2 := square(naturals)

	for r := range merge(res1, res2) {
		fmt.Println(r)
	}

}
