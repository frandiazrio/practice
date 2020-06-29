package main

import "fmt"

func main() {
	fmt.Println("Hello Señor")
	/*mp := make(map[int]int)
	mp[0] += +1
	mp[0] += 1
	fmt.Println(mp)*/

	for _, c := range "hello" {
		fmt.Println(string(c), 7/2)
	}
}
