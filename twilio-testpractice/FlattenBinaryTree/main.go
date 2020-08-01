package main

import "fmt"

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func FlattenBinaryTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if root.left == nil && root.right == nil {
		return root
	}

	left := FlattenBinaryTree(root.left)
	right := FlattenBinaryTree(root.right)

	if left == nil {
		return right
	}

	left.right = right
	root.right = left
	root.left = nil
	return left
}

func main() {
}
