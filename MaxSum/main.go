package main

import "fmt"
import "errors"
import "math"

func maxSum(arr []int) (int, error) {
	if len(arr) == 0 {
		return -1, errors.New("Empty array")
	}
	curr_sum := math.MinInt32
	max_sum := math.MinInt32

	for _, v := range arr {
		fmt.Println(max_sum)
		curr_sum = Max(v, curr_sum+v)
		max_sum = Max(max_sum, curr_sum)
	}

	return max_sum, nil
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	fmt.Println(maxSum([]int{1, 2, 3, -10, 2}))
}
