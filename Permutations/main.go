package main

import "fmt"

func permute(arr []int) [][]int {
	if len(arr) == 1 {
		return [][]int{arr}
	}

	res := [][]int{}
	for i, _ := range arr {
		take := arr[i]

		cp := make([]int, len(arr))
		copy(cp, arr)
		partial := append(cp[:i], cp[i+1:]...)
		for _, p := range permute(partial) {
			res = append(res, append(p, take))
		}

	}

	return res
}

func main() {

	fmt.Println(permute([]int{1, 2, 3, 4}))
}
