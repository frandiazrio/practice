package main

import (
	"fmt"
	"strings"
)

func Missing(s, t string) []string {
	//split words
	s_map := make(map[string]int)
	t_map := make(map[string]int)
	res := []string{}
	for _, c := range strings.Split(s, " ") {
		s_map[c] += 1
	}

	for _, c := range strings.Split(t, " ") {
		t_map[c] += 1
	}

	for k, seen := range s_map {
		_, ok := t_map[k]
		if !ok {
			for t := 0; t < seen; t++ {
				res = append(res, k)
			}
		}
	}
	return res
}

func main() {
	for _, word := range Missing("I am using HackerRank to improve programming", "am HackerRank to improve") {
		fmt.Println(word)
	}
}
