package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := make(chan string)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			str := s.Text()
			if len(str) == 0 {
				break
			} else {
				out <- str
			}
		}
		close(out)
	}()

	for c := range out {
		fmt.Println(c)
	}

}
