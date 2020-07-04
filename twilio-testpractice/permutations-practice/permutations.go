package main

import "fmt"

func permute(nums []int) [][]int {

	if len(nums) < 2 {
		return [][]int{nums}
	}

	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		cp := make([]int, len(nums))
		copy(cp, nums)

		take := nums[i : i+1]
		
		partial := append(cp[:i], cp[i+1:]...)
		fmt.Println(take, nums)
		for _, c := range permute(partial) {
			res = append(res, append(c, take...))
		}
	}

	return res
}

func main() {
	array1 := []int{1, 2, 3}
	//array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//array = append(array, 10)
	fmt.Println(array1)
	fmt.Println(permute(array1))
	
}
