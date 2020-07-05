package main

import (
	"fmt"
	"math"
)

func editDistance(str1, str2 string, m, n int)int{
	if m == 0{
		return n
	}

	if n == 0{
		return m
	}


	if str1[m-1] == str2[n-1]{
		return editDistance(str1, str2, m-1, n-1)
	}

	// Insert
	// Replace
	// Remove
	return 1 + Min(editDistance(str1, str2, m, n-1),editDistance(str1,str2,m-1, n-1),editDistance(str1, str2 , m-1, n))
}				


func Min(a ...int)int{
	var min int = math.MaxInt64

	for _, v := range a{
		if v < min{
			min = v
		}
	}

	return min
}

func main(){
	str1 := "sunday"
	str2 := "saturday"

	fmt.Println(editDistance(str1, str2, len(str1), len(str2)))
}