package main

import "fmt"

type pair struct {
	prefix string
	phone  string
}

func match(prefixes, numbers []string) []pair {
	//make a map of prefixes
	prefix_mp := make(map[string]struct{})

	// add them
	for _, p := range prefixes {
		prefix_mp[p] = struct{}{}
	}
	res := []pair{}

	//for every number
	for _, n := range numbers {
		pre := n

		for len(pre) != 0 {
			_, ok := prefix_mp[pre]
			if ok {
				new_pair := pair{pre, n}
				res = append(res, new_pair)
				break
			}

			pre = pre[:len(pre)-1]
		}
	}

	return res

}
func main() {
	prefixes := []string{"787", "939", "801", "787-21"}
	numbers := []string{"787-217-0452", "787-391-8639", "939-217-0875"}
	fmt.Println(match(prefixes, numbers))
}
