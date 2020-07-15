package main

import (
	"errors"
	"fmt"
)

func BinarySearch(arr []int, target int) (int, error) {
	lo := 0
	hi := len(arr) - 1

	return BinarySearchAux(arr, target, lo, hi)
}

func BinarySearchAux(arr []int, target int, low, hi int) (int, error) {

	if hi >= low {
		mid := low + (hi-low)/2

		if target < arr[mid] {
			return BinarySearchAux(arr, target, low, mid-1)
		} else if target > arr[mid] {
			return BinarySearchAux(arr, target, mid+1, hi)
		} else {
			return target, nil
		}
	} else {
		return -1, errors.New("Item not found")
	}
}

func main() {
	fmt.Println(BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))
	fmt.Println(BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11))

}
