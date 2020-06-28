package main

import "fmt"

func sumK(arr []int, k int)int{
	res := 0
	mp := make(map[int]int)
	sum := 0

	mp[0]=1 //we have already seen an initial sum of 0
	for _,v := range arr{
		sum += v
		_, ok := mp[sum-k]
		if ok {
			res++
		}
		mp[sum]++
	}

	return res

}

func main(){
	fmt.Println(sumK([]int{1,1,1, -1,2},2))
}