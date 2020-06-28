package main

import "fmt"

func check(v int)int{
	 if (v & 8) != 0{ 
		 return 1
	}
		return 0
}

func main(){
	fmt.Println(check(32))
	fmt.Println(check(77))
	fmt.Println(check(-1))
}