package main

import "fmt"

func powerset(set []string) [][]string {
	if len(set) == 0 {
		return [][]string{set}
	}

	res := [][]string{}
	take := set[0:1]
	partial := set[1:]

	for _, s := range powerset(partial) {
		res = append(res, s)
		res = append(res, append(s, take...))
	}

	return res

}

func main() {
	set := []string{"5","10", "11", "9", "5"}
	fmt.Println(powerset(set))

}
