package main

import "fmt"

// n = q*k + rem
// for a number to be divisible by k, rem = 0
// n = q1*k +rem1 + q2*k +rem2
// n = (q1-q2)*k + (rem1 - rem2)

// for it to be divisible rem1 must = rem2
// Every time we encounter more than one remainder we have a number that is divisible

func ksub(arr []int, k int) int {
	if len(arr) == 0 {
		return -1
	}

	karr := make([]int, k)
	// 0%k == 0
	karr[0] = 1
	sum := 0
	res := 0

	for _, v := range arr {
		sum += v
		sum %= k

		if sum < 0 {
			sum += k
		}

		res += karr[k]
		karr[k]++
	}

	return res
}

func main() {}
