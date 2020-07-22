// O(n)

package main

import "fmt"

/*
Let there be a subarray (i, j) whose sum is divisible by k
  sum(i, j) = sum(0, j) - sum(0, i-1)
Sum for any subarray can be written as q*k + rem where q
is a quotient and rem is remainder
Thus,
    sum(i, j) = (q1 * k + rem1) - (q2 * k + rem2)
    sum(i, j) = (q1 - q2)k + rem1-rem2

We see, for sum(i, j) i.e. for sum of any subarray to be
divisible by k, the RHS should also be divisible by k.
(q1 - q2)k is obviously divisible by k, for (rem1-rem2) to
follow the same, rem1 = rem2 where
    rem1 = Sum of subarray (0, j) % k
    rem2 = Sum of subarray (0, i-1) % k
*/

func ksub(arr []int, k int) int {
	kmap := make([]int, k)
	sum := 0 // current sum
	res := 0 //total results

	kmap[0] = 1 // A sum of 0 always divisible

	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
		sum %= k
		if sum < 0 {
			sum += k
		} // check for negative indices

		res += kmap[sum]

		kmap[sum]++

	}

	fmt.Println(kmap)

	return res

}

func main() {
	fmt.Println(ksub([]int{5, 10, 11, 9, 5}, 5))
	fmt.Println(ksub([]int{-5}, 5))
	fmt.Println(ksub([]int{7, 4, -10}, 5))
	fmt.Println(ksub([]int{4, 5, 0, -2, -3, 1}, 5))

	fmt.Println(-9 % 5)
}
