package main

import "fmt"

func BinarySearch(arr []int, lo, hi, x int) int {
	if lo <= hi {
		mid := lo + ((hi - lo) / 2)

		if x > arr[mid] {
			return BinarySearch(arr, mid+1, hi, x)
		} else if x < arr[mid] {
			return BinarySearch(arr, lo, mid, x)
		} else {
			return mid
		}

	}

	return -1
}

func KClosest(arr []int, k, x int) []int {
	index := BinarySearch(arr, 0, len(arr)-1, x)

	if len(arr) <= k {
		return arr
	}

	if index == -1 {
		return arr[:k+1]
	}

	i, j := index-1, index

	for (j - i + 1) != k {

		if i >= 0 && j < len(arr) {
			if (j-i+1)%2 == 0 {
				i--
			} else {
				j++
			}
		} else {
			if i == 0 {
				j++
			}

			if j == len(arr)-1 {
				i--
			}
		}

	}

	return arr[i : j+1]
}

func main() {
}
