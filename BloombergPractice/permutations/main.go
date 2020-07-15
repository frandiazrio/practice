package main

import "fmt"

func Permutations(arr []int) [][]int {
	if len(arr) < 1 {
		return [][]int{arr}
	}

	res := [][]int{}
	for i, v := range arr {
		cp := make([]int, len(arr))
		copy(cp, arr)
		partial := append(cp[:i], cp[i+1:]...)

		for _, p := range Permutations(partial) {
			res = append(res, append(p, v))
		}

	}

	return res
}

func main() {
	fmt.Println(Permutations([]int{1, 2, 3, 4}))
	fmt.Println([]int{})
}
