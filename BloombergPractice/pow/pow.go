package main

import "fmt"

func Pow(x float64, n int) float64 {
	if n < 0 {
		return 1 / Pow(x, -n)
	} else if n == 0 {
		return 1
	} else if n%2 == 0 {
		partial := Pow(x, n/2)
		return partial * partial
	}
	partial := Pow(x, n-1)
	return partial * x
}

func main() {
	fmt.Println(Pow(10.0, 2))
	fmt.Println(Pow(10.0, 5))
	fmt.Println(Pow(10.0, -1))
	fmt.Println(Pow(2, 32))
}
