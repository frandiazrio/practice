package main

import "fmt"

// P(S) = P(S-a) U (P(S-a)+a)
func powerSet(arr []int) [][]int {

	if len(arr) == 0 {
		return [][]int{arr}
	}

	take := arr[0]
	partial := arr[1:]
	res := [][]int{}
	for _, p := range powerSet(partial) {
		res = append(res, p)
		res = append(res, append(p, take))
	}

	return res
}
func main() {
	fmt.Println(powerSet([]int{1, 2, 3}))
}
