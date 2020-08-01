package main

import "fmt"

type BTreeNode struct {
	val               int
	left, right, next *BTreeNode
}

func helper(root *BTreeNode, level int, arr [][]*BTreeNode) [][]*BTreeNode {
	if root != nil {
		arr[level] = append(arr[level], root)
		arr = helper(root.left, level+1, arr)
		arr = helper(root.right, level+1, arr)
	}

	return arr

}
func levelOrderTraversal(root *BTreeNode) *BTreeNode {
	if root == nil {
		return nil
	}

	height := TreeHeight(root)

	arr := make([][]*BTreeNode, height)

	arr = helper(root, 0, arr)

	for lvl, _ := range arr {
		for i := 0; i < len(arr[lvl])-1; i++ {
			arr[lvl][i].next = arr[lvl][i+1]
		}

	}

	return root
}

func TreeHeight(root *BTreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + Max(TreeHeight(root.left), TreeHeight(root.right))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
}
