// C(n,r) = C(n-1,r-1) + C(n-1, r)

package main

import "fmt"

func Choose(n, k int64) int64 {
	if k > n {
		panic("Choose: k > n")
	}
	if k < 0 {
		panic("Choose: k < 0")
	}
	if n <= 1 || k == 0 || n == k {
		return 1
	}
	if newK := n - k; newK < k {
		k = newK
	}
	if k == 1 {
		return n
	}
	// Our return value, and this allows us to skip the first iteration.
	ret := int64(n - k + 1)
	for i, j := ret+1, int64(2); j <= k; i, j = i+1, j+1 {
		ret = ret * i / j
	}
	return ret
}

func main() {
	var n int64 = 5
	var k int64 = 0
	for k = 0; k <= n; k++ {
		fmt.Println(Choose(n, k))
	}
}
