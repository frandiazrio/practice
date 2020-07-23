package main

import "fmt"

func auxBinarySearch(arr []int, target, low, hi int) int {

	if low <= hi {
		mid := low + ((hi - low) / 2)

		if target < arr[mid] {
			return auxBinarySearch(arr, target, 0, mid)
		} else if target > arr[mid] {
			return auxBinarySearch(arr, target, mid+1, hi)
		} else if target == arr[mid] {
			return mid
		} else {
			return -1
		}
	}

	return -1
}

func BinarySearch(arr []int, target int) int {

	return auxBinarySearch(arr, target, 0, len(arr)-1)
}

func main() {
	fmt.Println()
}
