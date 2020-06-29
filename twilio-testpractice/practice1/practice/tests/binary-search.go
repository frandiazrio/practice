package main

import "fmt"

func binarySearch(array []int, target int, lowIndex int, highIndex int) (int, int) {
	if lowIndex > highIndex {
		return -1, -1
	}

	mid := lowIndex + (highIndex-lowIndex)/2

	if array[mid] < target {
		return binarySearch(array, target, mid+1, highIndex)
	} else if array[mid] > target {
		return binarySearch(array, target, lowIndex, mid)
	} else {
		return mid, target
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	//                 ^
	fmt.Println(binarySearch(arr, 3, 0, 6))
}
