package main

import "fmt"

func gen(nums ...int) <-chan int {

	out := make(chan int)

	go func() {
		for i := range nums {
			out <- i
		}

		close(out)
	}()

	return out

}

func cubed(input <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for i := range input {
			out <- i * i * i
		}

		close(out)
	}()

	return out
}

func printChannelOutput(input <-chan int) {
	for val := range input {
		fmt.Println(val)
	}
}

func main() {
	//generating numbers
	const num = 1000000
	arr := make([]int, num)
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	naturals := gen(arr...)
	res := cubed(naturals)
	printChannelOutput(res)

}
