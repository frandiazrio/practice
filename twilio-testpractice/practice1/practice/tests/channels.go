package main

import "fmt"

func greet(c chan string) {
	fmt.Println("Hello " + <-c)
	fmt.Println(<-c)
}

func main() {

	ch := make(chan string)
	defer close(ch)
	go greet(ch)

	ch <- "Francisco"

	ch <- "Hello"
}
