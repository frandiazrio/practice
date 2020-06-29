package main

import (
	"fmt"
	"math"
)

func AddStrings(s1, s2 string) int {
	atoi := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}

	i1 := StringToInt(atoi, s1)
	i2 := StringToInt(atoi, s2)
	return i1 + i2

}

func StringToInt(atoi map[string]int, s string) int {

	power := 0
	number := 0
	for i := len(s) - 1; i >= 0; i-- {
		number = number + (atoi[string(s[i])] * int(math.Pow10(power)))
		power++
	}

	return number
}

func main() {
	fmt.Println(AddStrings("200", "400"))

}
