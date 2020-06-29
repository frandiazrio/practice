package main

import "fmt"

func squares(ch chan int) {
	defer close(ch)
	for i := 0; i <= 100; i++ {
		ch <- i * i
	}

}

func main() {
	ch := make(chan int)
	go squares(ch)

	for {
		_, ok := <-ch
		fmt.Println(ok)
		if ok == false {
			break
		} else {
			//fmt.Println(val)
		}

	}
}
