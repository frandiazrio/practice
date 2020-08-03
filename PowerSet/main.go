package main

import "fmt"

func PowerSet(set []int) [][]int {
	if len(set) == 0 {
		return [][]int{set}
	}
	take := set[0]
	partial := set[1:]
	res := [][]int{}
	for _, s := range PowerSet(partial) {
		res = append(res, s)
		fmt.Println(s)
		res = append(res, append(s, take))
	}

	return res
}

func main() {

	fmt.Println(PowerSet([]int{1, 2, 3}))
}
