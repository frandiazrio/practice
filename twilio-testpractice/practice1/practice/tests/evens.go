package main

import "fmt"

func evens(arr []int) []int {
	res := []int{}
	for _, n := range arr {
		if n%2 == 0 {
			res = append(res, n)
		}
	}

	return res
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(evens(arr))

	n := "1223334455"

	c := n[:3]

	fmt.Println(c, n)
}
