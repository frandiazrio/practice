package main

import "fmt"

func permutations(array []int) [][]int {
	if len(array) < 2 {
		return [][]int{array}
	}

	res := [][]int{}
	for i := 0; i < len(array); i++ {
		take := array[i : i+1]
		cp := make([]int, len(array))
		copy(cp, array)

		partial := append(cp[:i], cp[i+1:]...)
		for _, v := range permutations(partial) {
			res = append(res, append(v, take...))
		}
	}

	return res
}

func main() {

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	array = append(array, 10)
	fmt.Println(array)
	fmt.Println([]int{})
	//fmt.Println(permutations(arrary))
}
