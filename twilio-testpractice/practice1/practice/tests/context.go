package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		fmt.Println("Canceling")
		cancel()
	}()

	sleepAndTalk(ctx, 5*time.Second, "Hello Context")
}

func sleepAndTalk(ctx context.Context, t time.Duration, str string) {
	time.Sleep(t)
	fmt.Println(str)
}
