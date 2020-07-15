package main

import "fmt"

func TwoSum(arr []int, k int) [][]int {
	// build map
	mp := make(map[int]int)
	for _, v := range arr {
		mp[v] = k - v
	}

	res := [][]int{}

	for _, v := range arr {
		comp, ok1 := mp[v]

		_, ok2 := mp[comp]
		if ok1 && ok2 {
			res = append(res, []int{v, comp})
			delete(mp, comp)
		}
	}

	return res
}
func main() {
	arr := []int{1, 2, 3, 7, 9, 11, -1}
	fmt.Println(TwoSum(arr, 10))
}
