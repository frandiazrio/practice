package main

import "fmt"
type Pair struct{
	a,b int
}
func diffPair(arr []int, k int) int {

	if len(arr) <= 1 || k < 0{
		return 0
	}


	res := make(map[Pair]struct{})
	mp := make(map[int]struct{})


	for _, number := range arr{
		_, ok1 := mp[number+k]
		_, ok2 := mp[number-k]
		
		if ok1 {
			res[Pair{min(number+k, number),max(number+k, number)}]=  struct{}{}
		}
		if ok2 {
			res[Pair{min(number-k, number),max(number-k, number)}] = struct{}{}
		}
		mp[number]=struct{}{}
	}
		

	return len(res)
}
func min(a,b int)int{
	if a < b {
		return a
	}

	return b
}

func max(a,b int)int{
	if a > b {
		return a
	}

	return b
}


func main() {
	fmt.Println(diffPair([]int{3, 1, 4, 1, 5}, 2))
	fmt.Println(diffPair([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(diffPair([]int{1, 3, 1, 5, 4}, 0))
}
