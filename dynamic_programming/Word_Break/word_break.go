package main

import "fmt"

func word_break(word string, mp map[string]struct{}) bool {

	dp := make([]bool, len(word)+1)

	for i := 1; i <= len(word); i++ {
		for j := 0; j < i; j++ {
			_, ok := mp[word[j:i]]

			if dp[j] && ok {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(word)]
}

func main() {
	mp := map[string]struct{}{
		"leet": struct{}{},
		"code": struct{}{},
	}

	fmt.Println(word_break("leetcode", mp))
}
