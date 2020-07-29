package main

import "math"

func maxSumSubArray(arr []int) int {
	if len(arr) < 1 {
		return -1
	}

	maxSum := 0
	currSum := arr[0]

	for i := 1; i < len(arr); i++ {
		currSum = int(math.Max(float64(arr[i]), float64(currSum+arr[i])))
		maxSum = int(math.Max(float64(currSum), float64(maxSum)))

	}

	return maxSum
}
