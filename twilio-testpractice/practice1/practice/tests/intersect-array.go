package main

import "fmt"

func Intersection(a, b []int) []int {
	// get min

	min, max := GetMinMax(a, b)
	// create map
	mp_max := make(map[int]int)
	mp_min := make(map[int]int)
	results := []int{}
	for _, v := range min {
		mp_min[v] += 1
	}

	for _, v := range max {
		mp_max[v] += 1
	}

	for k, count_max := range mp_max {

		count_min, ok_min := mp_min[k]

		if ok_min {
			minima := Minima(count_max, count_min)
			for i := 0; i < minima; i++ {
				results = append(results, k)
			}

			delete(mp_min, k)
		}

	}

	return results

}

func Minima(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func GetMinMax(a, b []int) ([]int, []int) {
	min := []int{}
	max := []int{}
	if len(a) < len(b) {
		min = a
		max = b

	} else {
		min = b
		max = a

	}

	return min, max
}

func main() {
	a := []int{1, 2, 3, 4, 3}
	b := []int{2, 3, 5}

	fmt.Println(Intersection(a, b))
	// intersection(a,b) -> {2,3}
}
