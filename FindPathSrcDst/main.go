package main

import "fmt"

type Node struct {
	val      int
	children []*Node
}

func helper(src, dst *Node, path []*Node, res [][]*Node) [][]*Node {
	head := path[len(path)-1]

	if head == dst {
		res = append(res, path)
	}
	for _, n := range head.children {
		path = append(path, n)
		helper(src, dst, path, res)
		path = path[:len(path)-1]
	}

	return res
}
func FindAllPaths(src, dst *Node) [][]*Node {
	if src == nil || dst == nil {
		return nil
	}

	path := []*Node{src}
	res := [][]*Node{}
	return helper(src, dst, path, res)
}

func main() {
}
