package main

import "fmt"

func power(x float64, n int) float64 {
	if n < 0 {
		power(x, -n)
	} else if n == 0 {
		return 1.0
	} else if n%2 == 0 {
		partial := power(x, n/2)
		return partial * partial
	}
	partial := power(x, (n-1)/2)
	return partial * partial * x

}

func main() {
	fmt.Println(power(3.25, 30))
}
