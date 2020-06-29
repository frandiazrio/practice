package main

import (
	"fmt"
	"sync"
)

type Node struct {
	key   int
	val   string
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

func (bst *BinarySearchTree) Insert(key int, val string) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	n := &Node{key, val, nil, nil}
	if bst.root == nil {
		bst.root = n
	} else {
		insertNode(bst.root, n)
	}

}

func insertNode(node, newNode *Node) {

	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}

	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}

}

func (bst *BinarySearchTree) PreOrder(f func(string)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	preOrder(bst.root, f)

}

func preOrder(node *Node, f func(string)) {
	if node != nil {
		f(node.val)
		preOrder(node.left, f)
		preOrder(node.right, f)
	}

}

func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		fmt.Printf(format+"%d\n", n.key)
		stringify(n.right, level)
	}
}

func main() {

	bst := BinarySearchTree{}
	bst.Insert(3, "3")
	bst.Insert(1, "1")
	bst.Insert(2, "2")
	bst.Insert(4, "4")

	stringify(bst.root, 0)
}
