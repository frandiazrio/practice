package main

import "fmt"

func Min(x, y, z int) int {
	min := x

	if y < min {
		min = y
	}

	if z < min {
		min = z
	}

	return min

}

func helper(str1, str2 string, m, n int) int {
	if m == 0 {
		return n
	}

	if n == 0 {
		return m
	}

	if str1[m] == str2[n] {
		return helper(str1, str2, m-1, n-1)
	}

	insert := helper(str1, str2, m-1, n)
	replace := helper(str1, str2, m-1, n-1)
	remove := helper(str1, str2, m, n-1)
	return 1 + Min(insert, remove, replace)

}

func EditDistance(str1, str2 string) int {
	return helper(str1, str2, len(str1), len(str2))
}

func main() {
}
