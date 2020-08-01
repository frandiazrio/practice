package main

import "fmt"

func main() {

	str := "Hello World"
	substr := str[0:5]
	fmt.Printf("%t: %v", substr, substr)
}
