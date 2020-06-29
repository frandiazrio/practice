package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func EvenLen(words string) (string, error) {
	if len(words) == 0 {
		return "", errors.New("wrong")
	}

	largest := ""
	for _, w := range strings.Split(words, " ") {
		if len(w) > len(largest) && len(w)%2 == 0 {
			largest = w
		}
	}

	return largest, nil
}

func main() {
	res, err := EvenLen("Slim shaddy is the best song ever")

	hello := "Hello"
	fmt.Println(hello[:2])
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(res)
	}

	for _, c := range hello {
		fmt.Println(string(c))
	}
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(19))

}
