package main

func Max(x, y, z int) int {
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

	// replace, insert, delete
	replace := helper(str1, str2, m-1, n-1)
	insert := helper(str1, str2, m-1, n)
	delete := helper(str1, str2, m, n-1)

	return Max(replace, insert, delete)
}

func longest_common_subsequence(str1, str2 string) int {

	return helper(str1, str2, len(str1)-1, len(str2)-1)
}

func main() {

}
