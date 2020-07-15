package main

import "sort"

func ThreeSum(arr []int, k int){
	sort.Slice(arr, func(i, j int)bool{
		return arr[i] < arr[j]
	})

	res := [][]int{} // i, j, k
	for i := 0; i< len(arr); i++{
		if arr[i]
	}
}

func main(){

}