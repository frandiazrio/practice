package main

import "fmt"

func kmp(str, pattern string) bool {
	table := buildPrefixSuffixTable(pattern)

	j := 0
	i := 0
	for i < len(str) {

		if j == len(pattern) {
			return true
		}

		if str[i] == pattern[j] {
			i++
			j++
		} else {
			j = table[j]
		}
	}

	return false

}

func buildPrefixSuffixTable(pattern string) []int {
	tb := make([]int, len(pattern))
	i, j := 0, 1

	for j < len(pattern) {
		if pattern[i] != pattern[j] {
			tb[j] = 0
			j++
		} else {
			tb[j] = tb[j-1] + 1
			j++
			i++
		}
	}
	return tb
}

func main() {
	fmt.Println(buildPrefixSuffixTable("dsgwadsgz"))
}
