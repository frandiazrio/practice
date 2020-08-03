package main

import "fmt"

func helper(str1, str2 string, m, n, count int) bool {

	if m == 0 && n == 0 && count == 1 {
		return true
	}

	if str1[m] == str2[n] && count <= 1 {
		return helper(str1, str2, m-1, n-1, count)
	} else {
		return false
	}

	insert := helper(str1, str2, m, n-1, count+1)
	delete := helper(str1, str2, m-1, n, count+1)
	replace := helper(str1, str2, m-1, n-1, count+1)

	return insert || delete || replace

}

func OneEditDistance(str1, str2 string) bool {
	return helper(str1, str2, len(str1)-1, len(str2)-1, 0)
}

func main() {
}
