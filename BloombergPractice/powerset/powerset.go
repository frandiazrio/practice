package main

import "fmt"

// P ({}) = {{}}
// P(S) = P(S-a) U (P(S-a)+a)
func PowerSet(S []string) [][]string {
	if len(S) == 0 {
		return [][]string{S}
	}

	res := [][]string{}
	a := S[0]

	partial := PowerSet(S[1:])

	for _, set := range partial {
		res = append(res, set)
		res = append(res, append(set, a))
	}

	return res

}

func main() {
	fmt.Println(PowerSet([]string{"a", "b", "c"}))
}
