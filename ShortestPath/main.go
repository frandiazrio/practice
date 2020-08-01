package main

import "fmt"
import "math"

type Node struct {
	val      int
	adjacent []*Node
}

func Dijkstra(src, dst *Node) []*Node {
	prev := make(map[*Node]*Node)
	dist := make(map[*Node]int)

	dist[src] = 0

}

func getPrev(mp map[*Node]*Node, node *Node) *Node {
	prev, ok := mp[node]

	if ok {
		return prev
	}

	return nil
}

func getDistance(mp map[*Node]int, node *Node) int {
	dist, ok := mp[node]

	if ok {
		return dist
	}

	return math.MaxInt32
}

func main() {
}
