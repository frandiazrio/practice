package main

import "fmt"

type pair struct{
	val int
	complement int
}

func distinct(arr []int, k int)[]pair{
	complement := make(map[int]int, len(arr))

	res := []pair{}
	for _,v := range arr{
		complement[v]=k-v
	}

	fmt.Println(complement)
	for _,v := range arr{
		v1, _ := complement[v]
		v2, ok2 := complement[v1]

		if ok2{
			res = append(res, pair{v1,v2})
		}
	}


	return res


}

func main(){
	fmt.Println(distinct([]int{1, 2, 3, 6, 7, 8, 9, 1}, 10))	
}
