package main

import (
	"fmt"
	"sort"
)

func main() {
	list := []int{1, 2, 3}
	fmt.Println(list)
	sort.Slice(list, func(i, j int) bool {
		return list[j] < list[i]
	})

	fmt.Println(list)
}
