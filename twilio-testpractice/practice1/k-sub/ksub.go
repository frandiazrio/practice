package main

import "fmt"

func ksub(arr []int, k int) [][]int {
	res := [][]int{}
	if len(arr) == 0{
		return res
	}

	if len(arr) == 1{
		if arr[0]%k ==0{
			res = append(res,arr)
		}
		return res
	}
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			s := sum(arr[i : j+1])
			fmt.Println("Arr ", arr[i : j+1])
			if s%k == 0 {
				res = append(res, arr[i:j+1])
			}
		}
	}

	return res
}

func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func main() {
	//fmt.Println(-10%5)
	fmt.Println(ksub([]int{5, 10, 11, 9, 5}, 5))
	fmt.Println(ksub([]int{-5}, 5))
	fmt.Println(ksub([]int{7,4,-10}, 5))
}
