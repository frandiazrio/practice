package main

func combinations(n, k int) [][]int {

	res := [][]int{}
	first := 1
	res = backTrack(first, k, n, []int{}, res)
	return res

}

func backTrack(first, k, n int, current []int, res [][]int) [][]int {
	if len(current) == k {
		res = append(res, current)

	}

	for i := first; i <= n; i++ {
		current = append(current, i)
		backTrack(i+1, k, n, current, res)
		current = current[:len(current)-1]
	}

	return res
}
