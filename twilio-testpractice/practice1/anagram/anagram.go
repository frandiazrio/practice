package main

import "fmt"

func anagrams(arr []string) [][]string {
	maps := []map[string]int{}

	for _, s := range arr {
		mp := wordToMap(s)

		if !mapExists(mp, maps) {
			maps = append(maps, mp)
		}
	}

	res := make([][]string, len(maps))
	for _, word := range arr {
		wmp := wordToMap(word)

		for i, mps := range maps {
			if sameMaps(wmp, mps) {
				res[i] = append(res[i], word)
			}
		}
	}

	return res

}

func wordToMap(s string) map[string]int {
	mp := make(map[string]int)
	for _, v := range s {
		mp[string(v)]++
	}

	return mp
}

func mapExists(mp map[string]int, maps []map[string]int) bool {

	for _, m := range maps {
		if sameMaps(mp, m) {
			return true
		}

	}
	return false
}
func sameMaps(mp1, mp2 map[string]int) bool {
	if len(mp1) == len(mp2) {
		for k1, v1 := range mp1 {
			v2, ok := mp2[k1]

			if !ok || v2 != v1 {
				return false
			}
		}
		return true
	}
	return false
}

func main() {
	fmt.Println("Checking for anagrams:")
	fmt.Println(anagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat", "hello"}))
}
