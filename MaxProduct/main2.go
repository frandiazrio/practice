package main

import "errors"

func MaxProduct(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("Empty array")
	}

	prev_min := arr[0]
	prev_max := arr[0]
	current_max := arr[0]
	current_min := arr[0]

	ans := arr[0]

	for i := 1; i < len(arr); i++ {
		current_max = Max(prev_max*arr[i], Min(prev_min*arr[i], arr[i]))

		current_min = Min(prev_min*arr[i], Max(prev_max*arr[i], arr[i]))

		ans = Max(current_max, ans)

		prev_min = current_min
		prev_max = current_max
	}

	return ans
}
