package main

import (
	"fmt"
	"project/linkedlist"
)

func main() {

	// some slice

	arr := []int{1, 2, 3, 4, 5, 6, 7}
	l := linkedlist.Init()
	l.Add("Hello")
	l.Add(2.5)
	l.Fill(arr)
	l.Print()
	node := l.GetIndex(2)
	fmt.Println(node.Val)
	l.Remove(node)
	l.Print()
	//fmt.F
	//ch := make(chan int)
}
