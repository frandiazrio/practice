package main

import "fmt"

func helper(arr []int, x, i, j int) int {

	if i <= j {
		mid := i + ((j - i) / 2)

		// Search upper
		if arr[mid] > arr[j] {
			return helper(arr, x, mid+1, j)
		} else if arr[mid] < arr[i] {
			return helper(arr, x, i, mid)
		} else {
			return mid
		}
	}

	return -1
}

func SearchRotated(arr []int, x int) int {

	if len(arr) == 0 {
		return -1
	}

	return helper(arr, x, 0, len(arr)-1)
}

func main() {
}
