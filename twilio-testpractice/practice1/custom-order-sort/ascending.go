package main

import (
	"fmt"
	"sort"
)

type KeyValue struct{
	key, val int
}

func main(){
	x := []int{1,2,2,1,-1,3,4,5,6,7,8,9,10,10}

	counter := make(map[int]int)
	for _,v := range x{
		counter[v]++
	}

	values := []KeyValue{}

	for k,v := range counter{
		values = append(values, KeyValue{k,v})
	}
	//fmt.Println(values)
	sort.Slice(values, func(i, j int)bool{
		return values[i].val < values[j].val
	})

	//fmt.Println(values)

	table := make(map[int][]int)

	for _,entry := range values{
		table[entry.val]= append(table[entry.val], entry.key)
	}

	fmt.Println(table)
	for k,_ := range table{
		sort.Slice(table[k], func(i, j int)bool{
			return table[k][i]< table[k][j]
		})
	}

	res := []int{}

	for k,t := range table{
		for _,e := range t{
			for i:=0; i<k; i++{
				res = append(res,e)
			}
		}

	}

	fmt.Println(res)




	
}